package gopl

import (
	"fmt"
	"strings"
)

type IntSet struct {
	set []uint64
}

func (i *IntSet) Has(x int) bool {
	word, bit := x/64, x%64
	return word<len(i.set) && ((1<<bit & i.set[word]) != 0)
}

func (i *IntSet) Add(x int)  {
	word, bit := x/64, x%64
	for word>=len(i.set) {
		i.set = append(i.set, 0)
	}

	i.set[word] |= 1<<bit
}

func (i *IntSet) Union(u *IntSet) {
	for j, us := range u.set {
		if j < len(i.set) {
			i.set[j] |= us
		}else {
			i.set = append(i.set, us)
		}
	}
}

func (i *IntSet) String() string {
	var buf []string
	for j, is := range i.set {
		if is == 0 {
			continue
		}
		for k := 0; k < 64; k++ {
			if is & (1 << k) != 0 {
				d := j*64 + k
				buf = append(buf, fmt.Sprintf("%d", d))
			}
		}
	}

	return "{" + strings.Join(buf, " ") + "}"
}

func (i *IntSet) Len() int {
	l := 0
	for _, is := range i.set {
		if is == 0 {
			continue
		}
		for k := 0; k < 64; k++ {
			if (is & (1<<k)) != 0 {
				l++
			}
		}
	}

	return l
}

func (i *IntSet) Remove(x int)  {
	word, bit := x/64, x%64
	i.set[word] = i.set[word] & (^(uint64(1)<<bit))
}

func (i *IntSet) Clear()  {
	i.set = []uint64{}
}
func (i *IntSet) Copy() *IntSet {
	x := make([]uint64, len(i.set))
	copy(x, i.set)
	return &IntSet{
		set: x,
	}
}

func (i *IntSet) AddAll(x ...int)  {
	for _, xd := range x {
		i.Add(xd)
	}
}

func (i *IntSet) IntersectWith(u *IntSet)  {
	for j, is := range i.set {
		if j >= len(u.set) {
			i.set[j] = 0
			continue
		}
		if is ==0 {
			continue
		}

		i.set[j] &= u.set[j]
	}
}

func (i *IntSet ) DifferentWith(u *IntSet)  {
	for j, is := range i.set {
		if j < len(u.set) {
			for k := 0; k < 64; k++ {
				d := is & u.set[j]
				i.set[j] = i.set[j] & ^d
			}
		}
	}
}
func (i *IntSet) SymmetricWith(u *IntSet) *IntSet {
	var arr []uint64
	for j, is := range i.set {
		if j < len(u.set) {
			arr = append(arr, is^u.set[j])
		}else {
			arr = append(arr, is)
		}
	}
	f := len(i.set)
	for f < len(u.set) {
		arr = append(arr, u.set[f])
		f++
	}

	return &IntSet{
		set: arr,
	}
}