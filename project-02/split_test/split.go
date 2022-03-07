package main

import (
	"strings"
)

func split(s string, tmp string) []string {
	var str []string
	index := strings.Index(s, tmp)
	for index > 0 {
		//append(str,s[:1] -> 0-i
		str = append(str, s[:index])
		//s[1+1:] 2->n
		s = s[index+1:]
		index = strings.Index(s, tmp)
	}
	str = append(str, s)
	return str
}
