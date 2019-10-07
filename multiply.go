package gomodmultiplies

import (
	"reflect"
)

// Multiply function for multiply numbers
func Multiply(items ...interface{}) int64 {
	res := int64(1)
	for _, item := range items {
		res *= convertToInt64(item)
	}
	return res
}

// MultiplyInt function for multiply numbers with integer params
func MultiplyInt(items ...int) int64 {
	res := 1
	for _, item := range items {
		res *= item
	}
	return int64(res)
}

// convertToInt64 is a function to convert numbers into interger 64
func convertToInt64(data interface{}) int64 {
	if data == nil {
		return 0
	}
	int64Type := reflect.TypeOf(int64(0))
	v := reflect.ValueOf(data)
	v = reflect.Indirect(v)
	if !v.Type().ConvertibleTo(int64Type) {
		return 0
	}
	res := v.Convert(int64Type)
	return res.Int()
}
