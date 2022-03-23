package util

import (
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
	t1map, _ := Struct2MapV2(&t1, true, false)
	fmt.Println(t1map)
	VerifyMapValueType(t1map)
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
