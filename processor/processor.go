package processor

import (
	"errors"
	"honey-badger/helper"
	"reflect"
)

// SniffableFunc
type SniffableFunc struct {
	Function interface{}
}

// SniffFunction prepares the parameters for the given function and calls the executor
func (f *SniffableFunc) SniffFunction() error {

	function := reflect.TypeOf(f.Function)
	paramsNumber := function.NumIn()
	in := make([]reflect.Value, paramsNumber)

	for i := 0; i < paramsNumber; i++ {
		param := function.In(i)
		if helper.ParamIsInt(param) {
			object := helper.GetRandInt(-100, 100)
			in[i] = reflect.ValueOf(object)
		}
	}
	if f.executeFunc(in) {
		return errors.New("panic found")
	}
	return nil
}

// executeFunc executes the given function and captures the panics
func (f *SniffableFunc) executeFunc(in []reflect.Value) (panics bool) {
	defer func() {
		if err := recover(); err != nil {
			panics = true
		}
	}()
	reflect.ValueOf(f.Function).Call(in)
	return panics
}

// ValidateInput validates the input
func (f *SniffableFunc) ValidateInput() error {
	if !helper.IsFunc(f.Function) {
		return errors.New("given param is not a function")
	}
	return nil
}
