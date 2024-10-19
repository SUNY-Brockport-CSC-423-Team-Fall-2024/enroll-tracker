package utils

import "reflect"

func IsZeroValueStruct(v interface{}) bool {
	return reflect.DeepEqual(v, reflect.Zero(reflect.TypeOf(v)).Interface())
}

func Contains[T comparable](val T, arr *[]T) bool {
	for _, v := range *arr {
		if val == v {
			return true
		}
	}
	return false
}
