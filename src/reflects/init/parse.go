package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

const file = "./.env"

var structName string

func main() {
	var config Config
	err := loadINI(file, &config)
	if err != nil {
		fmt.Printf("加载配置文件有误，%v\n", err)
	}

	fmt.Printf("%#v\n", config.MySQL)

	fmt.Printf("%#v\n", config.Redis)
}

// 加载配置文件
func loadINI(fileName string, data interface{}) (err error) {
	t := reflect.TypeOf(data)
	if t.Kind() != reflect.Ptr {
		err = fmt.Errorf("loadIni函数参数必须是一个指针")
		return
	}
	if t.Elem().Kind() != reflect.Struct {
		err = fmt.Errorf("loadIni函数参数必须是一个结构体指针")
		return
	}

	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}

	// string(b) // 将读取的字节类型的配置文件内容转换成字符串
	lineSlice := strings.Split(string(b), "\n")
	for _, line := range lineSlice { // 循环配置文件中的行
		line = strings.TrimSpace(line)                                    // 取出首尾空格
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") { // 如果是以;或#开头的行注释则忽略
			continue
		}

		if strings.HasPrefix(line, "[") || strings.HasSuffix(line, "]") { // 如果是[开头或]结尾，则是节
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 { // 如果是空行则忽略
				continue
			}

			// 根据sectionName节寻找对应的结构体
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				// fmt.Println(sectionName, field.Tag)
				if sectionName == field.Tag.Get("ini") {
					structName = field.Name
					break
				}
			}
		} else {
			// 只匹配有 = 的配置
			if strings.Contains(line, "=") {
				index := strings.Index(line, "=")
				key := strings.TrimSpace(line[:index])
				value := strings.TrimSpace(line[index+1:])
				v := reflect.ValueOf(data)
				sValue := v.Elem().FieldByName(structName) // 嵌套结构体值信息
				sType := sValue.Type()                     // 嵌套结构体类型信息
				if structObj := v.Elem().FieldByName(structName); structObj.Kind() != reflect.Struct {
					err = fmt.Errorf("data中的字段应该是一个结构体")
				} else {
					// 遍历嵌套结构体，赋值
					var fieldName string
					for i := 0; i < structObj.NumField(); i++ {
						field := sType.Field(i)

						if field.Tag.Get("ini") == key {
							fieldName = field.Name
							break
						}
					}

					fileObj := sValue.FieldByName(fieldName) // 取值
					// fmt.Println(fieldName, fileObj.Type().Kind())

					switch fileObj.Type().Kind() {
					case reflect.String:
						fileObj.SetString(value)
					case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
						v, _ := strconv.ParseInt(value, 10, 64) // 字符串转换成64位整型
						fileObj.SetInt(v)
					}

				}
			}
		}
	}

	return
}

type Config struct {
	MySQL `ini:"mysql"`
	Redis `ini:"redis"`
}

// MySQL 配置结构体
type MySQL struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

// Redis 配置结构体
type Redis struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
}
