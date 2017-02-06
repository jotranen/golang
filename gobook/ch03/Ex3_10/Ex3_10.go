package main

import (
	"fmt"
	"bytes"
	"unicode/utf8"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func comma2(s string) string {
	var buf bytes.Buffer

	charsCnt := utf8.RuneCountInString(s)
	prefix := charsCnt % 3
	sep := ","
	if prefix == 0 {
		sep = ""
	}

	buf.WriteString(s[0:prefix])

	for i := prefix; i < len(s); i = i+3 {
		buf.WriteString(sep)
		sep = ","
		buf.WriteString(s[i:i+3])
	}
	return buf.String()
}

func main() {
	res1 := comma("12345")
	res2 := comma2("12345678")

	fmt.Printf("res1: %s\n", res1)
	fmt.Printf("res2: %s\n", res2)
}
