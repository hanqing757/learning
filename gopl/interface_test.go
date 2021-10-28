package gopl

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"testing"
)

func TestWordAndLineCounter(t *testing.T) {
	in := "Spicy jalapeno pastrami ut ham turducken.\n Lorem sed ullamco, leberkas sint short loin strip steak ut shoulder shankle porchetta venison prosciutto turducken swine.\n Deserunt kevin frankfurter tongue aliqua incididunt tri-tip shank nostrud.\n"
	w := &WordAndLineCounter{}
	_, err := w.Write([]byte(in))
	if err != nil || w.wordCnt != 32 || w.lineCnt != 3 {
		t.Fatalf("Write err:%+v\n", err)
	}

	w1 := &WordAndLineCounter{}
	in1 := "Hello bytedance go\n"
	_,  err = fmt.Fprintf(w1, "%s", in1)
	if err != nil || w1.wordCnt != 3 || w1.lineCnt != 1 {
		t.Fatalf("Fprintf err: %+v\n", err)
	}
}

func TestTempReader(t *testing.T) {
	s := `<p>A: {{.A}}</p><p>B: {{.B}}</p>`
	tr, err := NewTempReader(s, TempParams{
		A: "<b>Hello!</b>",
		B: "<b>World!</b>",
	})
	if err != nil{
		t.Fatalf("NewTempReader err:%+v\n", err)
	}
	var res []byte
	var totalCnt int
	var n int
	for true {
		buf := make([]byte, 100)
		n , err = tr.Read(buf)
		if err == io.EOF {
			break
		}
		res = append(res, buf...)
		totalCnt += n
	}
	res = res[:totalCnt]  //需要按照实际copy出来的长度截取一下。否则buf包含本身初始化的len=100
	if string(res) != "<p>A: <b>Hello!</b></p><p>B: <b>World!</b></p>" {
		t.Fatalf("Read err")
	}
}

// 如果不自己for循环读的话可以使用ioutil.ReadAll来一次性获取
func TestTempReaderV2(t *testing.T) {
	s := `<p>A: {{.A}}</p><p>B: {{.B}}</p>`
	tr, err := NewTempReader(s, TempParams{
		A: "<b>Hello!</b>",
		B: "<b>World!</b>",
	})
	if err != nil{
		t.Fatalf("NewTempReader err:%+v\n", err)
	}
	str, err := ioutil.ReadAll(tr)
	if err != nil {
		t.Fatalf("ioutil.ReadAll err:%+v\n", err)
	}
	if string(str) != "<p>A: <b>Hello!</b></p><p>B: <b>World!</b></p>" {
		t.Fatalf("Read err")
	}
}

func TestLimitReader(t *testing.T) {
	r := strings.NewReader("some io.Reader stream to be read\n")
	lr := LimitReader(r, 14)
	b, err := ioutil.ReadAll(lr)
	if err != nil{
		t.Fatalf("ioutil.ReadAll err: %+v\n", err)
	}
	if string(b) !=  "some io.Reader" {
		t.Fatalf("error")
	}

	r1 := strings.NewReader("some io.Reader stream to be read\n")
	lr1 := LimitReader(r1, 14)
	var buf1 = make([]byte,20)
	n1, err := lr1.Read(buf1)
	buf1 = buf1[:n1]
	if string(buf1) != "some io.Reader" {
		t.Fatalf("error")
	}

	r2 := strings.NewReader("some io.Reader stream to be read\n")
	lr2 := LimitReader(r2, 14)
	var buf2 = make([]byte,10)
	_, _ = lr2.Read(buf2)
	if string(buf2) != "some io.Re" {
		t.Fatalf("error")
	}
	var buf3 = make([]byte,10)
	n3, _ := lr2.Read(buf3)
	buf3 = buf3[:n3]
	if string(buf3) != "ader" {
		t.Fatalf("error")
	}
}