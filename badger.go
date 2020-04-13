package badger

import (
	"honey-badger/processor"
)

//Sniff validates the input and starts the panic search
func Sniff(f interface{}) error {
	testable := processor.SniffableFunc{
		Function: f,
	}

	err := testable.ValidateInput()
	if err != nil {
		return err
	}

	err = testable.SniffFunction()
	if err != nil {
		return err
	}
	return nil
}
