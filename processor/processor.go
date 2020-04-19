package processor

import (
	"errors"
	"fmt"
	"github.com/LeoRodMrez/badger/constants"
	"github.com/LeoRodMrez/badger/helper"
	"reflect"
)

// SniffFunction prepares the parameters for the given function and calls the executor
func (f *Sniffer) SniffFunction() error {
	var execErr error
	v, err := f.GetSniffValuesSlice()
	if err != nil {
		return err
	}
	// temporal approach for POC | todo: implement recursively
	switch len(v) {
	case 1:
		execErr = f.Permutate1param(v)
	case 2:
		execErr = f.Permutate2params(v)
	case 3:
		execErr = f.Permutate3params(v)
	case 4:
		execErr = f.Permutate4params(v)
	case 5:
		execErr = f.Permutate5params(v)
	default:
		return errors.New(constants.SniffInvalidParamsError)
	}
	if execErr != nil {
		return execErr
	}
	return nil
}

// Permutate5params executes all possible options for 5 sets of values, one for each argument
func (f *Sniffer) Permutate5params(in [][]interface{}) error {
	n := 1
	for i := 0; i < len(in[0]); i++ {
		for j := 0; j < len(in[1]); j++ {
			for k := 0; k < len(in[2]); k++ {
				for l := 0; l < len(in[3]); l++ {
					for m := 0; m < len(in[4]); m++ {
						inParams := []interface{}{in[0][i], in[1][j], in[2][k], in[3][l], in[4][m]}
						refParams := helper.GetReflectValueParams(inParams)
						if f.executeFunc(refParams) {
							fmt.Printf("executions: %v\n", n)
							name := helper.GetFunctionName(f.Function)
							//todo: values are not visible for pointers -> translate inParams to legible values
							return errors.New(fmt.Sprintf("panic found in method %v for parameters %#v", name, inParams))
						}
						n++
					}
				}
			}
		}
	}
	fmt.Printf("executions: %v\n", n)
	return nil
}

// Permutate4params executes all possible options for 4 sets of values, one for each argument
func (f *Sniffer) Permutate4params(in [][]interface{}) error {
	n := 1
	for i := 0; i < len(in[0]); i++ {
		for j := 0; j < len(in[1]); j++ {
			for k := 0; k < len(in[2]); k++ {
				for l := 0; l < len(in[3]); l++ {
					inParams := []interface{}{in[0][i], in[1][j], in[2][k], in[3][l]}
					refParams := helper.GetReflectValueParams(inParams)
					if f.executeFunc(refParams) {
						fmt.Printf("executions: %v\n", n)
						name := helper.GetFunctionName(f.Function)
						return errors.New(fmt.Sprintf("panic found in method %v for parameters %#v", name, inParams))
					}
					n++
				}
			}
		}
	}
	fmt.Printf("executions: %v\n", n)
	return nil
}

// Permutate3params executes all possible options for 3 sets of values, one for each argument
func (f *Sniffer) Permutate3params(in [][]interface{}) error {
	n := 1
	for i := 0; i < len(in[0]); i++ {
		for j := 0; j < len(in[1]); j++ {
			for k := 0; k < len(in[2]); k++ {
				inParams := []interface{}{in[0][i], in[1][j], in[2][k]}
				refParams := helper.GetReflectValueParams(inParams)
				if f.executeFunc(refParams) {
					fmt.Printf("executions: %v\n", n)
					name := helper.GetFunctionName(f.Function)
					return errors.New(fmt.Sprintf("panic found in method %v for parameters %#v", name, inParams))
				}
				n++
			}
		}
	}
	fmt.Printf("executions: %v\n", n)
	return nil
}

// Permutate2params executes all possible options for 2 sets of values, one for each argument
func (f *Sniffer) Permutate2params(in [][]interface{}) error {
	n := 1
	for i := 0; i < len(in[0]); i++ {
		for j := 0; j < len(in[1]); j++ {
			inParams := []interface{}{in[0][i], in[1][j]}
			refParams := helper.GetReflectValueParams(inParams)
			if f.executeFunc(refParams) {
				fmt.Printf("executions: %v\n", n)
				name := helper.GetFunctionName(f.Function)
				return errors.New(fmt.Sprintf("panic found in method %v for parameters %#v", name, inParams))
			}
			n++
		}
	}
	fmt.Printf("executions: %v\n", n)
	return nil
}

// Permutate1param executes all possible options for 1 set of values, one for each argument
func (f *Sniffer) Permutate1param(in [][]interface{}) error {
	n := 1
	for i := 0; i < len(in[0]); i++ {
		inParams := []interface{}{in[0][i]}
		refParams := helper.GetReflectValueParams(inParams)
		if f.executeFunc(refParams) {
			fmt.Printf("executions: %v\n", n)
			name := helper.GetFunctionName(f.Function)
			return errors.New(fmt.Sprintf("panic found in method %v for parameters %#v", name, inParams))
		}
		n++
	}
	fmt.Printf("executions: %v\n", n)
	return nil
}

// GetSniffValuesSlice return a slice of slices containing the values to be executed for each parameter of the function
func (f *Sniffer) GetSniffValuesSlice() ([][]interface{}, error) {
	function := reflect.TypeOf(f.Function)
	paramsNumber := function.NumIn()
	in := make([][]interface{}, paramsNumber)

	for i := 0; i < paramsNumber; i++ {
		p, err := generateParam(function.In(i))
		if err != nil {
			return nil, err
		}
		in[i] = p
	}
	return in, nil
}

// generateParams returns a random parameter for the given type, returns an error if the kind is not supported
func generateParam(t reflect.Type) ([]interface{}, error) {

	switch t.Kind() {
	case reflect.Int:
		v := helper.GetIntSniffValues()
		b := make([]interface{}, len(v))
		for i := range v {
			b[i] = v[i]
		}
		return b, nil
	case reflect.Bool:
		v := helper.GetBoolSniffValues()
		b := make([]interface{}, len(v))
		for i := range v {
			b[i] = v[i]
		}
		return b, nil
	case reflect.String:
		v := helper.GetStringSniffValues()
		b := make([]interface{}, len(v))
		for i := range v {
			b[i] = v[i]
		}
		return b, nil
	case reflect.Ptr:
		switch t.Elem().Kind() {
		case reflect.Int:
			v := helper.GetIntPtrSniffValues()
			b := make([]interface{}, len(v))
			for i := range v {
				b[i] = v[i]
			}
			return b, nil
		case reflect.Bool:
			v := helper.GetBoolPtrSniffValues()
			b := make([]interface{}, len(v))
			for i := range v {
				b[i] = v[i]
			}
			return b, nil
		case reflect.String:
			v := helper.GetStringPtrSniffValues()
			b := make([]interface{}, len(v))
			for i := range v {
				b[i] = v[i]
			}
			return b, nil
		default:
			return nil, errors.New(fmt.Sprintf("param kind not supported: pointer to %v", t.Kind()))
		}
	default:
		return nil, errors.New(fmt.Sprintf("param kind not supported: %v", t.Kind()))
	}
}

// ValidateInput validates the input
func (f *Sniffer) ValidateInput() error {
	if !helper.IsFunc(f.Function) {
		return errors.New(constants.NotAFunctionError)
	}
	if helper.GetNumberOfParams(f.Function) > 5 || helper.GetNumberOfParams(f.Function) < 1 {
		return errors.New(constants.SniffInvalidParamsError)
	}
	return nil
}
