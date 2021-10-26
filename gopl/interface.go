package gopl

import (
	"bufio"
	"fmt"
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