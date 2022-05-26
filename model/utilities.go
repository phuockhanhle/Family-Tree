package model

import (
	"reflect"
)

func StructToMap(s interface{}) map[string]interface{} {
	values := reflect.ValueOf(s)
	struct_fields := reflect.TypeOf(s)

	mapFieldValue := make(map[string]interface{}, struct_fields.NumField())

	for i := 0; i < struct_fields.NumField(); i++ {
		mapFieldValue[struct_fields.Field(i).Name] = values.Field(i).Interface()
	}

	return mapFieldValue
}

func MapToStruct(m map[string]interface{}, s interface{}) interface{} {

	struct_field_values := reflect.New(reflect.TypeOf(s)).Elem()
	struct_field_values.Set(reflect.ValueOf(s))

	for k, v := range m {
		field := struct_field_values.FieldByName(k)
		if k == "Gender" {
			if v.(int64) == int64(Male) {
				v = Male
			} else if v.(int64) == int64(Female) {
				v = Female
			} else {
				panic("Unknown gender")
			}
		}
		field.Set(reflect.ValueOf(v))
	}

	return struct_field_values.Interface()
}
