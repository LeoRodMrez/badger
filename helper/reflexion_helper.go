package helper

import "reflect"

// IsFunc checks with reflexion that the given interface is a go function
func IsFunc(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Func
}

// ParamIsInt checks with reflexion that the given param is an integer
func ParamIsInt(t reflect.Type) bool {
	return t.Kind() == reflect.Int
}