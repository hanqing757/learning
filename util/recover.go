package util

import (
	"fmt"
	"runtime"
)

func Recover() {
	if r := recover(); r != nil {
		buf := make([]byte, 2048)
		n := runtime.Stack(buf, false)
		fmt.Printf("[recover panic]: err= %v, stack =  %s", r, buf[:n])
	}
}
