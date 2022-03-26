package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mohae/deepcopy"
	"reflect"
)

var (
	inputParamsErr = errors.New("input is not struct or ptr2struct")
)

func StructToMapWithJsonMashal(s interface{}) (map[string]interface{}, error) {
	st := reflect.TypeOf(s)
	if st.Kind() == reflect.Ptr {
		st = st.Elem()
	}
	if st.Kind() != reflect.Struct {
		return nil, inputParamsErr
	}

	data, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	res := map[string]interface{}{}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil

}

type FieldMappingHookFunc func(string) string

type ConvertConfig struct {
	//映射结构体字段到map key的hook；不为nil时优先使用
	FieldMappingHook FieldMappingHookFunc

	//在FieldMappingHook为nil时，使用TagName生成map key
	// TagName 默认为空，此时使用字段名
	TagName string

	//是否递归转换结构体为map
	IsRecursiveConv bool

	//map的interface value 转换为结构体字段时时候执行强类型转换
	IsStrictTypeConvert bool

	//未成功转换的字段
	// struct -> map 不可导出字段
	// map -> struct 在map中未找到字段或者赋值不成功
	Unused []string

	//转换过程的报错
	Errors error
}

type Converter struct {
	config *ConvertConfig
}

func NewConverter(conf *ConvertConfig) *Converter {
	return &Converter{
		config: conf,
	}
}

// StructToMap default
func StructToMap(input interface{}) (map[string]interface{}, error) {
	conf := &ConvertConfig{
		IsRecursiveConv: true,
	}
	converter := NewConverter(conf)
	return converter.StructToMap(input)
}

// MapToStruct  default
func MapToStruct(input interface{}, output interface{}) error {
	conf := &ConvertConfig{}
	converter := NewConverter(conf)
	return converter.MapToStruct(input, output)
}

func (c *Converter) GenMapKeyByField(field reflect.StructField) string {
	var mapKey string
	if c.config.FieldMappingHook != nil {
		mapKey = c.config.FieldMappingHook(field.Name)
	} else {
		if c.config.TagName != "" {
			mapKey = field.Tag.Get(c.config.TagName)
		}
	}
	if mapKey == "" {
		mapKey = field.Name
	}
	return mapKey
}

func (c *Converter) StructToMap(input interface{}) (map[string]interface{}, error) {
	defer Recover()

	st := reflect.TypeOf(input)
	sv := reflect.ValueOf(input)
	if st.Kind() == reflect.Ptr {
		st = st.Elem()
		sv = sv.Elem()
	}
	//非结构体或者空指针
	if st.Kind() != reflect.Struct || !sv.IsValid() {
		return nil, fmt.Errorf("input is not invalid struct")
	}
	res := map[string]interface{}{}
	for i := 0; i < st.NumField(); i++ {
		var (
			mapKey string
			mapVal interface{}
			err    error
		)

		fieldVal := sv.Field(i)
		if !fieldVal.CanInterface() { //不可导出字段不转换
			continue
		}

		mapKey = c.GenMapKeyByField(st.Field(i))

		fieldValInterface := fieldVal.Interface()
		if c.config.IsRecursiveConv && fieldVal.Kind() == reflect.Struct {
			mapVal, err = c.StructToMap(fieldValInterface)
			if err != nil {
				return nil, err
			}
		} else {
			mapVal = fieldValInterface
		}
		// deepcopy
		mapValCopy := deepcopy.Copy(mapVal)

		res[mapKey] = mapValCopy
	}
	return res, nil
}

func (c *Converter) MapToStruct(input interface{}, output interface{}) error {
	defer Recover()
	// 输入是map或者map指针
	mt := reflect.TypeOf(input)
	mv := reflect.ValueOf(input)
	if mt.Kind() == reflect.Ptr {
		mt = mt.Elem()
		mv = mv.Elem()
	}
	// 非map 或者是map 空指针
	if mt.Kind() != reflect.Map || !mv.IsValid() {
		return fmt.Errorf("input params is not valid map")
	}
	//map的必须要是string
	mapIter := mv.MapRange()
	for mapIter.Next() {
		k := mapIter.Key()
		if k.Kind() != reflect.String {
			return fmt.Errorf("map key is invalid")
		}
	}

	//output check
	ot := reflect.TypeOf(output)
	ov := reflect.ValueOf(output)
	if ot.Kind() != reflect.Ptr || ov.IsNil() {
		return fmt.Errorf("output is not ptr or nil ptr")
	}
	ott := ot.Elem()
	ovv := ov.Elem()
	// output 指向元素必须是结构体或者结构指针
	if ott.Kind() == reflect.Struct {
		//
	} else if ott.Kind() == reflect.Ptr && ott.Elem().Kind() == reflect.Struct {
		//如果是空指针，需要先初始化
		if ovv.IsNil() {
			tmpOvv := reflect.New(ott.Elem())
			ovv.Set(tmpOvv)
		}
		ott = ott.Elem()
		ovv = ovv.Elem()
	} else {
		return fmt.Errorf("output is not struct ot struct ptr")
	}

	for i := 0; i < ott.NumField(); i++ {
		fieldVal := ovv.Field(i)
		fieldType := ott.Field(i)
		// 可寻址 且是 结构体的可导出字段才可set
		if !fieldVal.CanSet() {
			fmt.Printf("Field %s can not be set\n", fieldType.Name)
		}
		mapKey := c.GenMapKeyByField(fieldType)
		mapVal := mv.MapIndex(reflect.ValueOf(mapKey))
		// key not found
		if mapVal.IsZero() {
			fmt.Printf("key:%s not found in map\n", mapKey)
			continue
		}

		//todo 实现弱类型赋值
		if mapVal.Elem().Type().AssignableTo(fieldType.Type) {
			fieldVal.Set(mapVal.Elem())
		} else {
			fmt.Printf("map key %s can not assignto field %s\n", mapKey, fieldType.Name)
		}

	}

	return nil
}
