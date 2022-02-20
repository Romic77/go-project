package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
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

	fmt.Printf("%#v\n", config)
}

func loadInit(fileName string, target interface{}) error {
	var err error

	t := reflect.TypeOf(target)
	// 如果种类是指针类型，那么t的元素
	if t.Kind() != reflect.Ptr {
		err = errors.New("data param should be a pointer")
		return err
	}

	//如果元素的类型不是结构体
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data param should be struct ")
		return err
	}

	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	var structName string
	lineStr := strings.Split(string(bytes), "\n")
	for index, line := range lineStr {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d syntax error", index+1)
				return err
			}

			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			//去除[]，获取里面的内容
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d synax error", index+1)
				return err
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
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line: %d syntax error%", index+1)
				return err
			}
			lines := strings.Split(line, "=")
			key := strings.TrimSpace(lines[0])
			value := strings.TrimSpace(lines[1])
			if value == "" {
				continue
			}

			v := reflect.ValueOf(target)
			//获得结构体的值
			sValue := v.Elem().FieldByName(structName)
			//获得结构体的类型 MysqlConfig/RedisConfig
			sType := sValue.Type()

			if sValue.Kind() != reflect.Struct {
				err = fmt.Errorf("target中的%s字段不是一个结构体", structName)
				return err
			}
			var fieldName string
			var fieldType reflect.StructField
			for i := 0; i < sValue.NumField(); i++ {
				field := sType.Field(i)
				fieldType = field
				if field.Tag.Get("ini") == key {
					//找到对应的字段
					fieldName = field.Name
					break
				}
			}

			fileObj := sValue.FieldByName(fieldName)

			switch fieldType.Type.Kind() {
			case reflect.String:
				fileObj.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				//字符串转int
				valueInt, _ := strconv.ParseInt(value, 10, 64)
				fileObj.SetInt(valueInt)
			case reflect.Bool:
				valueBool, _ := strconv.ParseBool(value)
				fileObj.SetBool(valueBool)
			case reflect.Float32, reflect.Float64:
				valueFloat, _ := strconv.ParseFloat(value, 64)
				fileObj.SetFloat(valueFloat)
			}
		}
	}
	return nil
}
