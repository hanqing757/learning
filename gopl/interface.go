package gopl

import (
	"bufio"
	"fmt"
	"io"
	"learning/util"
	"strings"
)

type WordAndLineCounter struct {
	wordCnt int
	lineCnt int
}

func (w *WordAndLineCounter) Write(p []byte) (cnt int, err error) {
	s := string(p)
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanWords)
	wc := 0
	for scanner.Scan() {
		wc++
	}
	if err = scanner.Err(); err != nil {
		fmt.Printf("ScanWords err:%+v\n", err)
		return
	}
	w.wordCnt = wc

	scanner = bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanLines)
	lc := 0
	for scanner.Scan() {
		lc++
	}
	if err = scanner.Err(); err != nil{
		fmt.Printf("ScanLines err %+v\n", err)
		return
	}
	w.lineCnt = lc

	cnt = wc + lc
	return 
}

func (w *WordAndLineCounter) String() string {
	str := fmt.Sprintf("word cnt:%d, line cnt:%d",w.wordCnt, w.lineCnt)

	return str
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	t := struct {
		io.Writer
	}{w}
	p := []byte{'a'}
	c, err := t.Write(p)
	if err != nil {
		fmt.Printf("Write err:%+v\n", err)
	}
	c64 := int64(c)
	return t, &c64
}

type Tree struct {
	value int
	left, right *Tree
}

// Sort 使用tree对s进行排序 先构建平衡二叉树，再中序遍历
func Sort(s []int) *Tree {
	var root *Tree
	for _, v := range s{
		root = Add(root, v)
	}
	return root
	//AppendVal(s[:0], root)
}

func AppendVal(s []int, r *Tree) []int {
	if r != nil{
		s = AppendVal(s, r.left)
		s = append(s, r.value)
		s = AppendVal(s, r.right)
	}
	return s
}

func Add(r *Tree, v int) *Tree {
	if r == nil {
		node := new(Tree)
		node.value = v
		return node
	}

	if v < r.value {
		d := Add(r.left, v)
		r.left = d
	}
	if v > r.value {
		d := Add(r.right, v)
		r.right = d
	}

	return r
}

func (t *Tree) String() string {
	var s []int
	s = AppendVal(s, t)
	str := util.Array2String(s, " ")
	return str
}