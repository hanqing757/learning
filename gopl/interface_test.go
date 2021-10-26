package gopl

import (
	"fmt"
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