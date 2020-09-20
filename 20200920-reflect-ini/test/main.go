package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

type MySQLConfig struct {
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
	MySQLConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadIni(filename string, data interface{}) (err error) {
	t := reflect.TypeOf(data)
	if t.Kind() != reflect.Ptr {
		err = errors.New("data param should be a pointer")
		return
	}

	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data param should be a struct")
		return
	}

	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	lineSlice := strings.Split(string(b), "\n")
	var structName string
	for idx, line := range lineSlice {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line: %d, syntax error", idx+1)
				return
			}
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line: %d, syntax error", idx+1)
				return
			}

			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					structName = field.Name
				}
			}
		} else {
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line: %d syntax error", idx+1)
				return
			}
			index := strings.Index(line, "=")
			key := strings.TrimSpace(line[:index])
			value := strings.TrimSpace(line[index+1:])
			v := reflect.ValueOf(data)
			sValue := v.Elem().FieldByName(structName)
			sType := sValue.Type()
			structObj := v.Elem().FieldByName(structName)
			if structObj.Kind() != reflect.Struct {
				err = fmt.Errorf("data 中的 %s 字段应该是一个结构体", structName)
				return
			}
			var fieldName string
			var fieldType reflect.StructField
			for i := 0; i < sValue.NumField(); i++ {
				field := sType.Field(i)
				fieldType = field
				if field.Tag.Get("ini") == key {
					fieldName = field.Name
					break
				}
			}
			if len(fieldName) == 0 {
				continue
			}
			fieldObj := sValue.FieldByName(fieldName)
			switch fieldType.Type.Kind() {
			case reflect.String:
				fieldObj.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				var valueInt int64
				valueInt, err = strconv.ParseInt(value, 10, 63)
				if err != nil {
					err = fmt.Errorf("line: %d, value type error", idx+1)
					return
				}
				fieldObj.SetInt(valueInt)
			case reflect.Bool:
				var valueBool bool
				valueBool, err = strconv.ParseBool(value)
				if err != nil {
					err = fmt.Errorf("line: %d, value type error", idx+1)
					return
				}
				fieldObj.SetBool(valueBool)
			case reflect.Float32, reflect.Float64:
				var valueFloat float64
				valueFloat, err = strconv.ParseFloat(value, 64)
				if err != nil {
					err = fmt.Errorf("line: %d, value type error", idx+1)
					return
				}
				fieldObj.SetFloat(valueFloat)
			}
		}

	}
	return
}
func main() {
	var cfg Config
	err := loadIni("./20200920-reflect-ini/test/test.ini", &cfg)
	if err != nil {
		fmt.Printf("load ini failed, err: %v", err)
		return
	}
	fmt.Printf("%#v\n", cfg)
}
