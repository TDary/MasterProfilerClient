package Tools

import (
	"encoding/json"
	"reflect"
)

func StructToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		data[field.Name] = v.Field(i).Interface()
	}
	return data
}

func MapToString(obj map[string]interface{}) string {
	dataType, _ := json.Marshal(obj)
	dataString := string(dataType)
	return dataString
}
