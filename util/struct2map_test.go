package util

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"reflect"
	"regexp"
	"strings"
	"testing"
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
	T2  `json:"x12"`
}

type T2 struct {
	Y1 int    `json:"y1"`
	Y2 string `json:"y2"`
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

func TestStructConv(t *testing.T) {
	//t1 := T1{
	//	X1:  111,
	//	X2:  222,
	//	X3:  333,
	//	X4:  444.4,
	//	X5:  555.5,
	//	X6:  false,
	//	X7:  "byte",
	//	X8:  []byte("dance"),
	//	X9:  []int{1, 2, 3},
	//	X10: []string{"hello", "world"},
	//	X11: map[int]int{
	//		888: 888,
	//	},
	//	T2: T2{
	//		Y1: 777,
	//		Y2: "nice",
	//	},
	//}
	//t1map, _ := Struct2MapWithJsonMashal(t1)
	//t1map, _ := Struct2MapV2(&t1, true, false)
	//fmt.Printf("%+v\n", t1map)
	//VerifyMapValueType(t1map)
}

func TestMapConv(t *testing.T) {
	// 对于map转结构体，数据可能是从JSON来的，value大概率不是[]int,T2这种确定类型
	//可能是 []interface{}，或者是map[string]interface{}
	//所以这类转换到结构体字段需要是弱类型转换，即[]interface{} 可以转为[]int或[]string, map[string]interface{} 可以转换为map[string]int或者struct等等
	//m := map[string]interface{}{
	//	"X1":  111,
	//	"X2":  int32(222),
	//	"X3":  int64(333),
	//	"X4":  float32(444.4),
	//	"X5":  555.5,
	//	"X6":  false,
	//	"X7":  "byte",
	//	"X8":  []byte("dance"),
	//	"X9":  []int{1, 2, 3},
	//	"X10": []string{"hello", "world"},
	//	"X11": map[int]int{
	//		888: 888,
	//	},
	//	"X12": T2{
	//		Y1: 777,
	//		Y2: "nice",
	//	},
	//}
	//var s1 T1
	//err := Map2Struct(m, &s1, false)
	//fmt.Printf("err:%+v, s1:%+v\n", err, s1)
	//
	//var s2 *T1
	//err = Map2Struct(m, s2, false)
	//fmt.Printf("err:%+v, s2:%+v", err, s2)
}

func TestMapStructureV1(t *testing.T) {
	t1 := T1{
		X1:  111,
		X2:  222,
		X3:  333,
		X4:  444.4,
		X5:  555.5,
		X6:  false,
		X7:  "byte",
		X8:  []byte("dance"),
		X9:  []int{1, 2, 3},
		X10: []string{"hello", "world"},
		X11: map[int]int{
			888: 888,
		},
		T2: T2{
			Y1: 777,
			Y2: "nice",
		},
	}
	var m1 = struct {
		X1 int
		X2 int64
	}{}
	err := mapstructure.Decode(t1, &m1)
	fmt.Printf("%+v, %+v\n", err, m1)
	//VerifyMapValueType(m1)
}

func TestMapstructureV2(t *testing.T) {
	m1 := map[string]interface{}{
		"name": 111, //转换时key的首字母大小写不敏感
		"x2":   222, //map的value 类型检测弱，int可以赋值给结构体的int32
		"X3":   int64(333),
		"X4":   float32(444.4),
		"X5":   555.5,
		"X6":   false,
		"X7":   "byte",
		"X8":   []byte("dance"),
		"X9":   []int{1, 2, 3},
		"X10":  []string{"hello", "world"},
		"X11": map[int]int{
			888: 888,
		},
		"X12": map[string]interface{}{
			"Y1": 777,
			"Y2": "nice",
		},
	}
	var t1 T1
	err := mapstructure.Decode(m1, &t1)
	fmt.Printf("%+v, %+v\n", err, t1)
}

func CamelCaseToSnake(s string) string {
	matchRe := regexp.MustCompile(`([a-z0-9])([A-Z])`) //匹配所有小写连接大小的位置
	r := matchRe.ReplaceAllString(s, "${1}_${2}")      //捕获替换匹配 sB 被 s_b替换
	return strings.ToLower(r)
}

func TestCamelCaseToSnake(t *testing.T) {
	s1 := "CalArrLen"
	s2 := "calArrLen"
	fmt.Println(CamelCaseToSnake(s1))
	fmt.Println(CamelCaseToSnake(s2))
}

func TestIS(t *testing.T) {
	//var x1 *T2
	//fmt.Println(reflect.ValueOf(x1).Elem().IsValid())
	//fmt.Println(reflect.ValueOf(x1).Elem().IsZero())
	//fmt.Println(reflect.ValueOf(x1).Elem().IsNil())

	var x2 *map[int]int

	//fmt.Println(reflect.ValueOf(x2).Elem().IsNil())
	fmt.Println(reflect.ValueOf(x2).Elem().IsValid())
	fmt.Println(reflect.ValueOf(x2).Elem().IsZero())

}

func TestCanSet(t *testing.T) {
	var x = struct {
		x1 int
		x2 string
	}{x1: 1, x2: "2"}
	xv := reflect.ValueOf(x)
	for i := 0; i < xv.NumField(); i++ {
		val := xv.Field(i)
		fmt.Println(val.CanSet())
	}
}
