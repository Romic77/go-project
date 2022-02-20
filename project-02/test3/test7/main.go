package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
)

type MysqlConfig struct {
	Address  string
	Port     int
	Username string
	Password string
}

func loadInit(c interface{}) {
	//打开文件
	file, err := os.Open("./mysql.ini")
	if err != nil {
		fmt.Printf("open file error,%v\n", err)
		return
	}
	//记得关闭文件
	defer file.Close()
	//换行读取读取文件内容
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("文件读完了")
			return
		}
		if err != nil {
			fmt.Printf("read from file failed,err:%v", err)
			return
		}
		line = strings.TrimSpace(line)

		if line == "[mysql]" {
			continue
		}
		//打印行信息
		fmt.Println(line)
		split := strings.Split(line, "=")
		fmt.Println(split[0], split[1])
		//reflectMysqlConfig(c, split)
	}
}

func reflectMysqlConfig(c interface{}, split []string) {
	//获得v的类型
	config := reflect.TypeOf(c)
	//指针类型获取真正type需要调用Elem -下面这段代码非常重要
	if config.Kind() == reflect.Ptr {
		config = config.Elem()
	}
	newConfig := reflect.New(config)
	if split[0] == "Address" {
		newConfig.Elem().FieldByName("Address").SetString(split[1])
		fmt.Println(newConfig.Elem().FieldByName("Address").String())
	}

	/*for i := 0; i < newConfig.Elem().NumField(); i++ {
		field := newConfig.Field(i)
		if (*split)[0] == "address" {
			field.Elem().FieldByName("address").SetString((*split)[1])
		}
	}*/
}

func main() {
	var config MysqlConfig
	str := []string{"Address", "127.0.0.1"}
	reflectMysqlConfig(&config, str)
	//loadInit(&config)
	//fmt.Println(config.Address, config.Port, config.Username, config.Password)
}
