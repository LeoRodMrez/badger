package badger

import (
	"github.com/LeoRodMrez/badger/processor"
)

//Sniff validates the input and starts the panic search
func Sniff(f interface{}) error {
	sniffable := processor.Sniffer{
		Function: f,
	}
	// input validation, sniffable.Function must be a go function
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
