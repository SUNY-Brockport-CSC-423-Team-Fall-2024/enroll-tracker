package utils

import "reflect"

func IsZeroValueStruct(v interface{}) bool {
	return reflect.DeepEqual(v, reflect.Zero(reflect.TypeOf(v)).Interface())
}

type Argon2IDParams struct {
	memorySize  int
	iterations  int
	parallelism int
}
