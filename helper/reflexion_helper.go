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

// GetReflectValueParams takes a slice of interfaces and returns the equivalent slice of reflect values
func GetReflectValueParams(in []interface{}) []reflect.Value {
	r := make([]reflect.Value, len(in))
	for i, v := range in {
		r[i] = reflect.ValueOf(v)
	}
	return r
}

// GetNumberOfParams returns the number of params for the given function
func GetNumberOfParams(f interface{}) int {
	function := reflect.TypeOf(f)
	return function.NumIn()
}
