package gopl

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"html/template"
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

type TempReader struct {
	s string
	idx int
}

type TempParams struct {
	A template.HTML
	B template.HTML
}

func NewTempReader(s string, tp TempParams) (*TempReader, error) {
	var buf bytes.Buffer
	temp := template.Must(template.New("escape").Parse(s))
	if err := temp.Execute(&buf, tp); err != nil{
		fmt.Printf("NewTempReader err: %+v\n", err)
		return nil, err
	}

	return &TempReader{
		s: buf.String(),
		idx: 0,
	},nil
}

func (t *TempReader) Read(b []byte) (n int, err error)  {
	if t.idx >= len(t.s) {
		return 0, io.EOF
	}

	l := copy(b, t.s[t.idx:])
	t.idx += l
	return l, nil
}

// LimitedReader 实现LimitReader
type LimitedReader struct {
	R io.Reader
	N int64   //max bytes remaining
}

func (l *LimitedReader) Read(b []byte) (n int, err error)  {
	if l.N <= 0 {
		return 0, io.EOF
	}

	n1, err := l.R.Read(b)
	if err == io.EOF {
		return 0, err
	}
	if int64(n1) <= l.N {
		l.N = l.N - int64(n1)
		return n1, nil
	}else {
		b = b[:l.N]
		l.N = 0
		return len(b), nil
	}
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitedReader{
		R: r,
		N: n,
	}
}

type Celsius float64
type Fahrenheit float64
type Kelvim float64

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9)}
func KToC(k Kelvim) Celsius {return Celsius(k-273.15)}

func (c Celsius) String() string {
	return fmt.Sprintf("%fC",c)
}

type celsiusFlag struct {
	Celsius
}

func (f *celsiusFlag) Set(s string) error {
	var val float64
	var unit string
	fmt.Sscanf(s, "%f%s",&val, &unit)
	switch unit {
	case "C":
		f.Celsius = Celsius(val)
		return nil
	case "F":
		f.Celsius = FToC(Fahrenheit(val))
		return nil
	case "K":
		f.Celsius = KToC(Kelvim(val))
		return nil
	default:
		return fmt.Errorf("invalid tempurature %q", s)
	}
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}



