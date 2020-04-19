package processor

import (
	"errors"
	"fmt"
	"github.com/LeoRodMrez/badger/constants"
	"reflect"

	"github.com/LeoRodMrez/badger/helper"
)

// Sniffer encapsulates the function
type Sniffer struct {
	Function interface{}
}

// BruteForceSniffFunction prepares the parameters for the given function and calls the executor
func (f *Sniffer) BruteForceSniffFunction() error {
	//random search
	for i := 0; i < 10000; i++ {
		inParams, err := f.buildBruteForceParams()
		if err != nil {
			return err
		}
		refParams := helper.GetReflectValueParams(inParams)
		if f.executeFunc(refParams) {
			name := helper.GetFunctionName(f.Function)
			return errors.New(fmt.Sprintf("panic found in method %v for parameters %#v", name, inParams))
		}
	}
	return nil
}

// buildBruteForceParams builds the slice with parameters for further execution
func (f *Sniffer) buildBruteForceParams() ([]interface{}, error) {
	function := reflect.TypeOf(f.Function)
	paramsNumber := function.NumIn()
	in := make([]interface{}, paramsNumber)

	for i := 0; i < paramsNumber; i++ {
		p, err := generateBruteForceParam(function.In(i))
		if err != nil {
			return nil, err
		}
		in[i] = p
	}
	return in, nil
}

// generateParams returns a random parameter for the given type, returns an error if the kind is not supported
func generateBruteForceParam(t reflect.Type) (interface{}, error) {
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
		value, err = generateBruteForceValueForPointer(t.Elem())
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New(fmt.Sprintf("param kind not supported: %v", t.Kind()))
	}

	return value, nil
}

// generateBruteForceValueForPointer generates a random pointer for a given reflect type
func generateBruteForceValueForPointer(t reflect.Type) (interface{}, error) {
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

// ValidateBruteForceInput validates the input
func (f *Sniffer) ValidateBruteForceInput() error {
	if !helper.IsFunc(f.Function) {
		return errors.New(constants.NotAFunctionError)
	}
	if helper.GetNumberOfParams(f.Function) < 1{
		return errors.New(constants.BruteForceSniffInvalidParamsError)
	}
	return nil
}
