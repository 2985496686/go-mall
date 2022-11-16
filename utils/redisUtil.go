package utils

import (
	"errors"
	"reflect"
)

func StructToMap(Object interface{}) (m map[string]interface{}, e error) {
	ObjectType := reflect.TypeOf(Object)
	ObjectValue := reflect.ValueOf(Object)
	if ObjectType.Kind() == reflect.Ptr {
		ObjectType = ObjectType.Elem()
	}
	if ObjectType.Kind() != reflect.Struct {
		return nil, errors.New("StructToMap only access struct or structPtr to Map")
	}
	m = make(map[string]interface{})
	for i := 0; i < ObjectType.NumField(); i++ {
		fieldType := ObjectType.Field(i)
		fieldValue := ObjectValue.Field(i)
		if tag := fieldType.Tag.Get("toMap"); tag != "" {
			m[tag] = fieldValue.Interface()
		}
	}
	return
}
