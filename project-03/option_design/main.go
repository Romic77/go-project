package main

import "fmt"

type Options struct {
	strOption1 string
	strOption2 string
	strOption3 string
	intOption1 int
	intOption2 int
	intOption3 int
}

type Option func(opts *Options)

func InitOptions1(opts ...Option) {
	options := &Options{}
	for _, opt := range opts {
		//调用函数，在函数里，给穿进去
		opt(options)
	}
	fmt.Printf("init options %#v\n", options)
}

func withStrOption1(str string) Option {
	return func(opts *Options) {
		opts.strOption1 = str
	}
}

func withStrOption2(str string) Option {
	return func(opts *Options) {
		opts.strOption1 = str
	}
}

func main() {
	InitOptions1(withStrOption1("str1"), withStrOption2("str2"))
}
