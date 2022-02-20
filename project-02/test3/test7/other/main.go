package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
}

type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func main() {
	var config Config
	/*str := []string{"Address", "127.0.0.1"}
	reflectMysqlConfig(&config, str)*/
	loadInit("./config.ini", &config)

	//fmt.Println(config.Address, config.Port, config.Username, config.Password)
}

func loadInit(fileName string, target interface{}) {
	var err error

	t := reflect.TypeOf(target)
	// 如果种类是指针类型，那么t的元素
	if t.Kind() != reflect.Ptr {
		err = errors.New("data param should be a pointer")
		return
	}

	//如果元素的类型不是结构体
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data param should be struct ")
		return
	}

	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	var structName string
	lineStr := strings.Split(string(bytes), "\n")
	for index, line := range lineStr {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d syntax error", index+1)
				return
			}

			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			//去除[]，获取里面的内容
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d synax error", index+1)
				return
			}
			//反射就是用来找到变量名称和变量值
			//reflect.Typeof() 变量名称
			//reflect.Valueof() 变量的值
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				//说明找到了对应的嵌套结构体
				if sectionName == field.Tag.Get("ini") {
					//filed.Name 获取变量名称
					structName = field.Name
					fmt.Printf("找到%s对应的结构体%s\n", sectionName, structName)
				}
			}
		} else {

		}
	}

}
