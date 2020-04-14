package processor

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/LeoRodMrez/badger/helper"
)

// Sniffer encapsulates the function
type Sniffer struct {
	Function interface{}
}

// SniffFunction prepares the parameters for the given function and calls the executor
func (f *Sniffer) SniffFunction() error {
	//random search
	for i := 0; i < 1000; i++ {
		inParams, err := f.buildParams()
		if err != nil {
			panic(err.Error())
		}
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
func (f *Sniffer) buildParams() ([]interface{}, error) {
	function := reflect.TypeOf(f.Function)
	paramsNumber := function.NumIn()
	in := make([]interface{}, paramsNumber)

	for i := 0; i < paramsNumber; i++ {
		p, err := generateParam(function.In(i))
		if err != nil {
			return nil, err
		}
		in[i] = p
	}
	return in, nil
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

func generateParam(t reflect.Type) (interface{}, error) {
	var value interface{}
	var err error

	switch t.Kind() {
	case reflect.Int:
		value = helper.GetRandInt(-1000, 1000)
	case reflect.Bool:
		value = helper.GetRandBool()
	case reflect.String:
		value = helper.GetRandString()
	case reflect.Ptr:
		value, err = generateValueForPointer(t.Elem())
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New(fmt.Sprintf("param kind not supported: %v", t.Kind()))
	}

	return value, nil
}

// generateValueForPointer generates a random pointer for a given reflect type
func generateValueForPointer(t reflect.Type) (interface{}, error) {
	var value interface{}

	switch t.Kind() {
	case reflect.Int:
		value = helper.GetRandomIntPtr(-1000,1000)
	case reflect.Bool:
		value = helper.GetRandBoolPtr()
	case reflect.String:
		value = helper.GetRandStringPtr()
	default:
		return nil, errors.New(fmt.Sprintf("param kind not supported: pointer to %v", t.Kind()))
	}

	return value, nil
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
