package util

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func InStringSlice(stack []string, needle string) bool {
	for _, s := range stack {
		if s == needle {
			return true
		}
	}

	return false
}

func Array2String(s []int, delim string) string {
	var buf bytes.Buffer
	for i, v := range s {
		buf.WriteString(strconv.Itoa(v))
		if i < len(s)-1 {
			buf.WriteString(delim)
		}
	}

	return buf.String()
}

func Array2StringV2(s []int, delim string) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(s)), delim), "[]")
}
