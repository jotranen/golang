package main

import (
	"bytes"
	"unicode/utf8"
	"fmt"
	"strings"
)

/*
 * Exercise 3.11: Enhance comma so that it deals correctly with floating-point numbers and an optional sign.
 */

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

func comma3(s string) string {
	sign := ""
	s1 := s
	if strings.Index(s, "-") == 0 {
		sign = "-"
		s1 = s[1:]
	}

	i := strings.Index(s1, ".")
	fmt.Println("Index: ", i)
	if i > -1 {
		fmt.Printf("Passing: %s\n", s1[:i])
		return sign + comma2(s1[:i]) + s1[i:]
	} else {
		return sign + comma2(s1)
	}
}

func main() {
	res1 := comma("-123.45")
	fmt.Printf("res1: %s\n", res1)
	res2 := comma2("12345678")
	fmt.Printf("res2: %s\n", res2)

	res3 := comma3("-123456.78")
	fmt.Printf("res3: %s\n", res3)
}
