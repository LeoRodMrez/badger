package main

import (
	"fmt"
	"strings"

	"github.com/LeoRodMrez/badger"
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
	if a == 0 && strings.Contains(b,"j") && !c && strings.Contains(b,"%") {
		panic("")
	}
	return true
}
