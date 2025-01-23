package djan_go

import (
	"encoding/json"
	"reflect"
	"strings"
)

func GetObjectSchemaJson[T any](obj T) (string, error) {
	mp := GetTypeMap(obj)
	jsonString, err := json.Marshal(mp)
	if err != nil {
		return "", err
	}

	return string(jsonString), nil
}

func GetTypeMap[T any](obj T) map[string]interface{} {
	mp := GetTypeMapObject(reflect.TypeOf(obj))
	return mp
}

func GetTypeMapObject(t reflect.Type) map[string]interface{} {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	mp := make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := field.Name
		ftype := field.Type

		many := ""

		if field.Type.Kind() == reflect.Slice {

			ftype = field.Type.Elem()
			many = "many:"
		}
		ftypestring := ftype.String()
		if strings.Contains(ftypestring, "()") {
			continue
		}
		if ftypestring == "string" || ftypestring == "bool" {
			mp[value] = many + ftypestring
			continue
		}
		if ftypestring == "float32" || ftypestring == "float64" {
			ftypestring = "float"
			mp[value] = many + ftypestring
			continue
		}

		inttypes := []string{
			"int", "int8", "int16", "int32", "int64",
			"uint", "uint8", "uint16", "uint32", "uint64", "uintptr", "byte", "rune",
		}

		for _, inttype := range inttypes {
			if ftypestring == inttype {
				ftypestring = "int"
				mp[value] = ftypestring
				break
				continue
			}
		}

		if ftype.Kind() == reflect.Struct {
			mp[many+value] = GetTypeMapObject(ftype)
		} else {
			mp[value] = many + ftypestring
		}

	}

	return mp
}
