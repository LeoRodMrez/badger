package main

import (
	"badger"
	"fmt"
)

func main() {
	fmt.Println("Hi I'm the Honey Badger, lets find some panics!")

	err := badger.Sniff(myFunc)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("NOTHING")
	}
}

func myFunc(a int, b string, c bool, j string) bool {
	if a == 0 && b == "" && !c && j == "" {
		panic("")
	}
	return true
}
