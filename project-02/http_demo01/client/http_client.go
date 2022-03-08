package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	//f1()
	f2()
}

//构造GET请求
func f2() {
	data := url.Values{}
	data.Set("username", "f2")
	urlObj, _ := url.Parse("http://localhost:9090/golang")
	urlObj.RawQuery = data.Encode()
	reader := strings.NewReader("sb")
	request, _ := http.NewRequest("GET", urlObj.String(), reader)

	//禁用disableKeepAlives
	disableKeepAlives := http.Transport{DisableKeepAlives: true}
	client := http.Client{
		Transport: &disableKeepAlives,
	}
	response, _ := client.Do(request)
	defer response.Body.Close()
}

func f1() {
	response, err := http.Get("http://localhost:9090/golang?username=zhangsan")
	if err != nil {
		return
	}

	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	fmt.Println(string(result))
}
