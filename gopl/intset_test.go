package gopl

import "testing"

func TestIntSet(t *testing.T) {
	i1 := &IntSet{}
	i1.AddAll([]int{5,6,7,8,9}...)
	if !i1.Has(7) {
		t.Fatal("Has fail")
	}
	if i1.Has(2) {
		t.Fatal("Has fail")
	}
	i1.Remove(9)
	str := i1.String()
	if str != "{5 6 7 8}" {
		t.Fatal("String fail")
	}

	i2 := &IntSet{}
	i2.AddAll([]int{7,8,9,10}...)
	i1.Union(i2)
	if i1.Len() != 6 || !i1.Has(9) {
		t.Fatal("Union fail")
	}

	i1.Clear()
	i1.AddAll([]int{5,6,7,8}...)
	i2.Clear()
	i2.AddAll([]int{7,8,9,10}...)
	i1.IntersectWith(i2)
	if i1.String() != "{7 8}" {
		t.Fatal("IntersectWith err")
	}
	i1.Clear()
	i1.AddAll([]int{5,6,7,8}...)
	i2.Clear()
	i2.AddAll([]int{7,8,9,10}...)
	i1.DifferentWith(i2)
	if i1.String() != "{5 6}" {
		t.Fatal("DifferentWith err")
	}

	i1.Clear()
	i1.AddAll([]int{5,6,7,8}...)
	i2.Clear()
	i2.AddAll([]int{7,8,9,10}...)
	i3 := i1.SymmetricWith(i2)
	if i3.String() != "{5 6 9 10}" {
		t.Fatal("SymmetricWith err")
	}
}