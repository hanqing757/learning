package util

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestStructConv(t *testing.T) {
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
		X12: T2{
			Y1: 777,
			Y2: "nice",
		},
	}
	//t1map, _ := Struct2MapV1(t1)
	t1map, _ := Struct2MapV2(&t1, true, true)
	fmt.Println(t1map)
	VerifyStructFieldType(t1map)
}

func TestMapConv(t *testing.T) {
	m := map[string]interface{}{
		"X1":  111,
		"X2":  int32(222),
		"X3":  int64(333),
		"X4":  float32(444.4),
		"X5":  555.5,
		"X6":  false,
		"X7":  "byte",
		"X8":  []byte("dance"),
		"X9":  []int{1, 2, 3},
		"X10": []string{"hello", "world"},
		"X11": map[int]int{
			888: 888,
		},
		"X12": T2{
			Y1: 777,
			Y2: "nice",
		},
	}
	var s1 T1
	err := Map2Struct(m, &s1, false)
	fmt.Printf("err:%+v, s1:%+v\n", err, s1)

	var s2 *T1
	err = Map2Struct(m, &s2, false)
	fmt.Printf("err:%+v, s2:%+v", err, s2)
}

func Test1(t *testing.T) {
	var x1 []int
	var x2 = make([]int, 0)
	s1, _ := json.Marshal(x1)
	fmt.Println(string(s1), s1)
	s2, _ := json.Marshal(x2)
	fmt.Println(string(s2), s2)
}

type T11 struct {
	X1 int    `json:"x1"`
	X2 string `json:"x2"`
}

type T12 struct {
	X1 int
	x2 string
}

func Test2(t *testing.T) {
	t11 := T11{
		X1: 1,
		X2: "222",
	}
	s, _ := json.Marshal(t11)
	var t12 *T12
	fmt.Println(string(s))
	_ = json.Unmarshal(s, &t12)
	fmt.Println(t12)
}
