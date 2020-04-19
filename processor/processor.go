package processor

import (
	"errors"
	"github.com/LeoRodMrez/badger/constants"
	"github.com/LeoRodMrez/badger/helper"
)

// SniffFunction prepares the parameters for the given function and calls the executor
func (f *Sniffer) SniffFunction() error {

	n := helper.GetNumberOfParams(f.Function)
	// temporal approach for POC | todo: implement recursively
	switch n {
	case 1:

	case 2:

	case 3:

	case 4:

	case 5:

	default:
		return errors.New(constants.SniffInvalidParamsError)
	}
	return nil
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

func permute5params(in [][]interface{}) {

}
