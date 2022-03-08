package main

import (
	"strings"
)

func split(s string, sep string) []string {
	var str []string
	index := strings.Index(s, sep)
	for index > 0 {
		//append(str,s[:1] -> 0-i
		str = append(str, s[:index])
		//s[1+1:] 2->n
		s = s[index+len(sep):]
		index = strings.Index(s, sep)
	}
	str = append(str, s)
	return str
}
