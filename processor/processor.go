package processor

import (
	"errors"

	"github.com/LeoRodMrez/badger/helper"
)

// SniffFunction prepares the parameters for the given function and calls the executor
func (f *Sniffer) SniffFunction() error {
	_, err := helper.GetSniffNumberOfExecutions(f.Function)
	if err != nil {
		return err
	}
	return nil
}

// ValidateInput validates the input
func (f *Sniffer) ValidateInput() error {
	if !helper.IsFunc(f.Function) {
		return errors.New("given param is not a function")
	}
	if helper.GetNumberOfParams(f.Function) > 5 {
		return errors.New("max number of parameters supported is 5")
	}
	return nil
}