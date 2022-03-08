package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
) //http请求协议的框架封装，已经很完善

func f1(response http.ResponseWriter, request *http.Request) {
	all, _ := ioutil.ReadAll(request.Body)
	queryParams := request.URL.Query()

	fmt.Println(request.URL)
	fmt.Println(request.URL.Query())
	fmt.Println(queryParams.Get("username"))
	fmt.Println(request.Method)
	fmt.Println(string(all))
	response.Write([]byte("hello world!"))
}

func f2(response http.ResponseWriter, request *http.Request) {
	all, _ := ioutil.ReadAll(request.Body)
	fmt.Println(request.URL)
	fmt.Println(request.Method)
	fmt.Println(string(all))
	response.Write([]byte("hello f2"))
}

//监听端口为9090，访问url为golang，然后进入函数f1处理response
func main() {
	http.HandleFunc("/golang", f1)
	http.HandleFunc("/f2", f2)
	http.ListenAndServe("0.0.0.0:9090", nil)
}
