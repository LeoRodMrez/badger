package processor

import (
	"badger/helper"
	"errors"
	"fmt"
	"reflect"
)

// Sniffer encapsulates the function
type Sniffer struct {
	Function interface{}
}

// SniffFunction prepares the parameters for the given function and calls the executor
func (f *Sniffer) SniffFunction() error {
	//random search
	for i := 0; i < 1000; i++ {
		inParams := f.buildParams()
		refParams := getReflectValueParams(inParams)
		if f.executeFunc(refParams) {
			name := helper.GetFunctionName(f.Function)
			return errors.New(fmt.Sprintf("panic found in method %v for parameters %#v", name, inParams))
		}
	}
	return nil
}

// getReflectValuesParams takes a slice of interfaces and returns the equivalent slice of reflect values
func getReflectValueParams(in []interface{}) []reflect.Value {
	r := make([]reflect.Value, len(in))
	for i, v := range in {
		r[i] = reflect.ValueOf(v)
	}
	return r
}

// buildParams builds the slice with parameters for further execution
func (f *Sniffer) buildParams() []interface{} {
	function := reflect.TypeOf(f.Function)
	paramsNumber := function.NumIn()
	in := make([]interface{}, paramsNumber)

	for i := 0; i < paramsNumber; i++ {
		in[i] = generateParam(function.In(i))
	}
	return in
}

// buildZeroValuesParams builds the slice with zero value parameters for further execution
func (f *Sniffer) buildZeroValuesParams() []interface{} {
	function := reflect.TypeOf(f.Function)
	paramsNumber := function.NumIn()
	in := make([]interface{}, paramsNumber)

	for i := 0; i < paramsNumber; i++ {
		in[i] = reflect.Zero(function.In(i))
	}
	return in
}

func generateParam(t reflect.Type) interface{} {
	var value interface{}

	switch t.Kind() {
	case reflect.Int:
		value = helper.GetRandInt(-1000, 1000)
	case reflect.Bool:
		value = helper.GetRandBool()
	case reflect.String:
		value = helper.GetRandString(15)
	}

	return value
}

// executeFunc executes the given function and captures the panic
func (f *Sniffer) executeFunc(in []reflect.Value) (panic bool) {
	defer func() {
		if err := recover(); err != nil {
			panic = true
		}
	}()
	reflect.ValueOf(f.Function).Call(in)
	return panic
}

// ValidateInput validates the input
func (f *Sniffer) ValidateInput() error {
	if !helper.IsFunc(f.Function) {
		return errors.New("given param is not a function")
	}
	return nil
}
