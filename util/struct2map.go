package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

type T1 struct {
	X1  int         `json:"x1"`
	X2  int32       `json:"x2"`
	X3  int64       `json:"x3"`
	X4  float32     `json:"x4"`
	X5  float64     `json:"x5"`
	X6  bool        `json:"x6"`
	X7  string      `json:"x7"`
	X8  []byte      `json:"x8"`
	X9  []int       `json:"x9"`
	X10 []string    `json:"x10"`
	X11 map[int]int `json:"x11"`
	X12 T2          `json:"x12"`
}
type T2 struct {
	Y1 int    `json:"y1"`
	Y2 string `json:"y2"`
}

var (
	inputParamsErr = errors.New("input is not struct or ptr2struct")
)

func Struct2MapV1(s interface{}) (map[string]interface{}, error) {
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

func assertInt(k string, v interface{}) {
	if _, ok := v.(int); ok {
		fmt.Printf("conv success, key:%s, expect type:int, got:%T; val:%v\n", k, v, v)
	} else {
		fmt.Printf("conv failed, key:%s, expect type:int, got:%T; val:%v\n", k, v, v)
	}
}

func assertInt32(k string, v interface{}) {
	if _, ok := v.(int32); ok {
		fmt.Printf("conv success, key:%s, expect type:int32, got:%T; val:%v\n", k, v, v)
	} else {
		fmt.Printf("conv failed, key:%s, expect type:int32, got:%T; val:%v\n", k, v, v)
	}
}
func assertInt64(k string, v interface{}) {
	if _, ok := v.(int64); ok {
		fmt.Printf("conv success, key:%s, expect type:int64, got:%T; val:%v\n", k, v, v)
	} else {
		fmt.Printf("conv failed, key:%s, expect type:int64, got:%T; val:%v\n", k, v, v)
	}
}

func assertFloat32(k string, v interface{}) {
	if _, ok := v.(float32); ok {
		fmt.Printf("conv success, key:%s,expect type:float32, got:%T;  val:%v\n", k, v, v)
	} else {
		fmt.Printf("conv failed, key:%s, expect type:float32, got:%T; val:%v\n", k, v, v)
	}
}
func assertFloat64(k string, v interface{}) {
	if _, ok := v.(float64); ok {
		fmt.Printf("conv success, key:%s, expect type:float64, got:%T; val:%v\n", k, v, v)
	} else {
		fmt.Printf("conv failed, key:%s, expect type:float64, got:%T; val:%v\n", k, v, v)
	}
}

func assertBool(k string, v interface{}) {
	if _, ok := v.(bool); ok {
		fmt.Printf("conv success, key:%s,expect type:bool, got:%T; val:%v\n", k, v, v)
	} else {
		fmt.Printf("conv failed, key:%s, expect type:bool, got:%T; val:%v\n", k, v, v)
	}
}

func assertString(k string, v interface{}) {
	if _, ok := v.(string); ok {
		fmt.Printf("conv success, key:%s, expect type:string, got:%T; val:%v\n", k, v, v)
	} else {
		fmt.Printf("conv failed, key:%s, expect type:string, got:%T; val:%v\n", k, v, v)
	}
}

func assertByteSlice(k string, v interface{}) {
	if _, ok := v.([]byte); ok {
		fmt.Printf("conv success, key:%s,expect type:[]byte, got:%T; val:%v\n", k, v, v)
	} else {
		fmt.Printf("conv failed, key:%s, expect type:[]byte, got:%T; val:%v\n", k, v, v)
	}
}

func assertIntSlice(k string, v interface{}) {
	if _, ok := v.([]int); ok {
		fmt.Printf("conv success, key:%s,expect type:[]int, got:%T; val:%v\n", k, v, v)
	} else {
		fmt.Printf("conv failed, key:%s, expect type:[]int, got:%T; val:%v\n", k, v, v)
	}
}

func assertStringSlice(k string, v interface{}) {
	if _, ok := v.([]string); ok {
		fmt.Printf("conv success, key:%s,expect type:[]string, got:%T; val:%v\n", k, v, v)
	} else {
		fmt.Printf("conv failed, key:%s, expect type:[]string, got:%T; val:%v\n", k, v, v)
	}
}
func assertMapIntInt(k string, v interface{}) {
	if _, ok := v.(map[int]int); ok {
		fmt.Printf("conv success, key:%s,expect type:map[int]int, got:%T; val:%v\n", k, v, v)
	} else {
		fmt.Printf("conv failed, key:%s, expect type:map[int]int, got:%T; val:%v\n", k, v, v)
	}
}

func assertT2(k string, v interface{}) {
	if _, ok := v.(T2); ok {
		fmt.Printf("conv success, key:%s,expect type:T2, got:%T; val:%v\n", k, v, v)
	} else {
		fmt.Printf("conv failed, key:%s, expect type:T2, got:%T; val:%v\n", k, v, v)
	}
}

func VerifyMapValueType(m map[string]interface{}) {
	if m == nil {
		fmt.Println("input map is nil")
	}
	for k, v := range m {
		switch k {
		case "X1", "x1":
			assertInt(k, v)
		case "X2", "x2":
			assertInt32(k, v)
		case "X3", "x3":
			assertInt64(k, v)
		case "X4", "x4":
			assertFloat32(k, v)
		case "X5", "x5":
			assertFloat64(k, v)
		case "X6", "x6":
			assertBool(k, v)
		case "X7", "x7":
			assertString(k, v)
		case "X8", "x8":
			assertByteSlice(k, v)
		case "X9", "x9":
			assertIntSlice(k, v)
		case "X10", "x10":
			assertStringSlice(k, v)
		case "X11", "x11":
			assertMapIntInt(k, v)
		case "X12", "x12":
			assertT2(k, v)
		}
	}

}

// Struct2MapV2 struct 转换为map[string]interface{}
// s必须是结构体或者结构体指针
func Struct2MapV2(s interface{}, useJsonTag bool, isRecursiveConvSuct bool) (map[string]interface{}, error) {
	defer Recover()

	st := reflect.TypeOf(s)
	sv := reflect.ValueOf(s)
	if st.Kind() == reflect.Ptr {
		st = st.Elem()
		sv = sv.Elem()
	}
	if st.Kind() != reflect.Struct {
		return nil, fmt.Errorf("input is not struct, is %+v", st.Kind())
	}
	res := map[string]interface{}{}
	for i := 0; i < st.NumField(); i++ {
		var (
			mapKey string
			mapVal interface{}
			err    error
		)

		if useJsonTag {
			mapKey = st.Field(i).Tag.Get("json")
		} else {
			mapKey = st.Field(i).Name
		}
		fieldVal := sv.Field(i)
		fieldValInterface := fieldVal.Interface()
		if isRecursiveConvSuct && fieldVal.Kind() == reflect.Struct {
			mapVal, err = Struct2MapV2(fieldValInterface, useJsonTag, isRecursiveConvSuct)
			if err != nil {
				return nil, err
			}
		} else {
			mapVal = fieldValInterface
		}
		res[mapKey] = mapVal
	}
	return res, nil
}

// Map2Struct map结构转换为结构体
// m是map[string]interface{}或者map[string]interface{}的指针
// output需要是（结构体或者结构体指针）的指针
func Map2Struct(m interface{}, output interface{}, useJsonTag bool) error {
	defer Recover()
	// 输入是map或者map指针
	mt := reflect.TypeOf(m)
	mv := reflect.ValueOf(m)
	if mt.Kind() == reflect.Ptr {
		mt = mt.Elem()
		mv = mv.Elem()
	}
	if mt.Kind() != reflect.Map {
		return fmt.Errorf("input params is not map or prt2map")
	}
	//map的必须要是map[string]interface{}
	mapIter := mv.MapRange()
	for mapIter.Next() {
		k := mapIter.Key()
		v := mapIter.Value()
		if k.Kind() != reflect.String || v.Kind() != reflect.Interface {
			return fmt.Errorf("map key or value not invalid")
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
		// 那些不可以set?
		if !fieldVal.CanSet() {
			fmt.Printf("Field %s can not be set\n", fieldType.Name)
		}
		var mapKey string
		if useJsonTag {
			mapKey = fieldType.Tag.Get("json")
		} else {
			mapKey = fieldType.Name
		}
		mapVal := mv.MapIndex(reflect.ValueOf(mapKey))
		// key not found
		if mapVal.IsZero() {
			fmt.Printf("key:%s not found in map\n", mapKey)
			continue
		}

		if mapVal.Elem().Type().AssignableTo(fieldType.Type) {
			fieldVal.Set(mapVal.Elem())
		} else {
			fmt.Printf("map key %s can not assignto field %s\n", mapKey, fieldType.Name)
		}

	}

	return nil
}
