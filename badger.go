package badger

import (
	"github.com/LeoRodMrez/badger/processor"
)

//Sniff validates the input and starts the quick panic search
func Sniff(f interface{}) error {
	sniffable := processor.Sniffer{
		Function: f,
	}
	// input validation
	err := sniffable.ValidateInput()
	if err != nil {
		return err
	}

	// start searching for panics
	err = sniffable.SniffFunction()
	if err != nil {
		return err
	}
	return nil
}

//BruteForceSniff validates the input and starts the panic search with brute force executions
func BruteForceSniff(f interface{}) error {
	sniffable := processor.Sniffer{
		Function: f,
	}
	// input validation
	err := sniffable.ValidateBruteForceInput()
	if err != nil {
		return err
	}

	// start searching for panics
	err = sniffable.BruteForceSniffFunction()
	if err != nil {
		return err
	}
	return nil
}
