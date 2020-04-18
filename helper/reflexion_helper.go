package helper

import (
	"errors"
	"fmt"
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

// GetSniffNumberOfExecutions returns the number of executions needed for a given function based on the parameters kind
func GetSniffNumberOfExecutions(f interface{}) (int, error) {
	n := 1
	function := reflect.TypeOf(f)
	paramsNumber := function.NumIn()

	for i := 0; i < paramsNumber; i++ {
		switch function.In(i).Kind() {
		case reflect.Int:
			n *= 5
		case reflect.Bool:
			n *= 2
		case reflect.String:
			n *= 5
		case reflect.Ptr:
			switch function.In(i).Elem().Kind() {
			case reflect.Int:
				n *= 6
			case reflect.Bool:
				n *= 3
			case reflect.String:
				n *= 6
			default:
				return 0, errors.New(fmt.Sprintf("param kind not supported: pointer to %v", function.In(i).Elem().Kind()))
			}
		default:
			return 0, errors.New(fmt.Sprintf("param kind not supported: %v", function.In(i).Kind()))
		}
	}

	return n, nil
}