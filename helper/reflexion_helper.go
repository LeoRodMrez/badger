package helper

import (
	"reflect"
	"runtime"
)

// IsFunc checks with reflexion that the given interface is a go function
func IsFunc(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Func
}

// GetFunctionName returns the function name
func GetFunctionName(v interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(v).Pointer()).Name()
}