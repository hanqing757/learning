package gopl

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"io"
	"learning/util"
	"log"
	"os"
	"reflect"
	"sort"
	"strings"
	"text/tabwriter"
	"time"
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

type Track struct {
	Title string
	Artist string
	Album string
	Year int
	Length time.Duration
}

var Tracks = []*Track{
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	t, err := time.ParseDuration(s)
	if err != nil {
		panic("ParseDuration err")
	}
	return t
}

// printTacks
func PrintTacks(tracks []*Track) {
	format := "%v\t%v\t%v\t%v\t%v\t\n"
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(w, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(w, format, "-----", "------", "-----", "----", "------")
	for _, track := range tracks {
		fmt.Fprintf(w, format, track.Title, track.Artist, track.Album, track.Year, track.Length)
	}
	w.Flush()
}

type ByDefault []*Track

func (b ByDefault) Len() int {
	return len(b)
}
func (b ByDefault) Less(i, j int) bool {
	if b[i].Title != b[j].Title {
		return b[i].Title < b[j].Title
	}
	if b[i].Year != b[j].Year {
		return b[i].Year < b[j].Year
	}
	if b[i].Length != b[j].Length {
		return b[i].Length < b[j].Length
	}

	return false
}
func (b ByDefault) Swap(i, j int)  {
	b[i], b[j] = b[j], b[i]
}

// CustomSort 实现sort.Interface的类型不一定是slice，也可以是结构体 用结构体实现ByDefault的排序实现
type CustomSort struct {
	t []*Track
	less func(x, y *Track) bool
}

func (c CustomSort) Len() int {
	return len(c.t)
}
func (c CustomSort) Less(i, j int) bool  {
	return c.less(c.t[i], c.t[j])
}
func (c CustomSort) Swap(i, j int)  {
	c.t[i], c.t[j] = c.t[j], c.t[i]
}

func MyCustomSort() {
	sort.Sort(CustomSort{Tracks, func(x, y *Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		if x.Length != y.Length {
			return x.Length < y.Length
		}

		return false
	}})
	PrintTacks(Tracks)
}

// MyCustomSortV2 更灵活的自定义排序
func MyCustomSortV2(column string, increase bool) {
	var less func(x, y *Track) bool
	switch column {
	case "Title":
		less = func(x, y *Track) bool {
			return x.Title < y.Title
		}
	case "Year":
		less = func(x, y *Track) bool {
			return x.Year < y.Year
		}
	case "Length":
		less = func(x, y *Track) bool {
			return x.Length < y.Length
		}
	default:
		panic("colume name err")
	}
	if increase {
		sort.Sort(CustomSort{Tracks, less})
	}else {
		sort.Sort(sort.Reverse(CustomSort{Tracks, less}))
	}
	PrintTacks(Tracks)
}

type MultiSort struct {
	t []*Track
	rule []string  //保存点击状态。倒序进行排序判定
}

func (m MultiSort) Len() int {
	return len(m.t)
}
func (m MultiSort) Less(i, j int) bool {
	for k := len(m.rule)-1; k > -1; k-- {
		switch m.rule[k] {
		case "Title":
			if m.t[i].Title != m.t[j].Title {
				return m.t[i].Title < m.t[j].Title
			}
		case "Year":
			if m.t[i].Year != m.t[j].Year {
				return m.t[i].Year < m.t[j].Year
			}
		case "Length":
			if m.t[i].Length != m.t[j].Length {
				return m.t[i].Length < m.t[j].Length
			}
		default:
			panic("column name err")
		}
	}
	return true
}
func (m MultiSort) Swap(i, j int)  {
	m.t[i], m.t[j] = m.t[j], m.t[i]
}

// IsPalindrome 通过sort.Interface实现回文判断
func IsPalindrome(s sort.Interface) bool {
	defer func() {
		if p := recover(); p != nil{
			log.Fatalf("err: %+v", p)
		}
	}()
	if s == nil || reflect.ValueOf(s).IsNil() {
		panic("nil err")
	}

	leng := s.Len()
	for i := 0; i < leng/2; i++ {
		if !s.Less(i, leng-1-i) && !s.Less(leng-1-i, i) {
			continue
		}else {
			return false
		}
	}
	return true
}