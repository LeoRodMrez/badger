package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hi I'm the Honey Badger, lets find some panics!")
	/*err := badger.Sniff(sum)
	if err != nil {
		fmt.Printf("error: %v", err)
	}*/
}


/*
func sum(a, b int) (int, int) {
	fmt.Printf("\nadding %v to %v equals %v", a, b, a+b)
	panic("\nim panicking")
	return a + b, a
}*/
