package utils

import (
	"fmt"
	"reflect"
	"strings"
)

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// StructToMap 使用反射将结构体转换为 map[string]interface{}
func StructToMap(obj interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	// 获取对象的反射值
	v := reflect.ValueOf(obj)

	// 如果是指针，取实际指向的值
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// 只处理结构体
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("input must be a struct or a pointer to struct")
	}

	// 获取结构体类型信息（字段名、tag等）
	//t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		value := v.Field(i)

		// 获取 json tag，如果没有则使用字段名
		tag := field.Tag.Get("json")
		if tag == "" {
			tag = strings.ToLower(field.Name)
		}

		// 忽略空值字段（可选）
		if isEmptyValue(value) {
			continue
		}

		result[tag] = value.Interface()
	}

	return result, nil
}

// 判断值是否为空（0、""、false、nil 等）
func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.String:
		return v.String() == ""
	case reflect.Bool:
		return !v.Bool()
	case reflect.Ptr, reflect.Interface, reflect.Slice, reflect.Map, reflect.Chan:
		return v.IsNil()
	}
	return false
}
