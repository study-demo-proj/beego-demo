package tool

import (
	"reflect"
)

/**
目前只适用于修改字段都是字符串且不填时传空字符和数字不填时传0的情况
 */
func BuildUpdateParams(input interface{}) map[string]interface{} {
	getType := reflect.TypeOf(input)
	getType = getType.Elem()
	result := make(map[string]interface{})

	getValue := reflect.ValueOf(input)
	getValue = getValue.Elem()

	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		//fieldType := reflect.TypeOf(value)
		//fmt.Println(fieldType.Name(),"======", "=======", )
		if value != nil && value != "" &&value != 0 {
			result[field.Name] = value
		}
	}
	return result
}
