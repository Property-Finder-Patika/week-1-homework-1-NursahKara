package main

import (
	"fmt"
)

func main() {
	var name string = "Nursah"
	var greeting = createGreet(name) // function assigned to variable greeting
	fmt.Printf("%s", greeting)
}

func createGreet(name string) string { // Function defined with return type string
	greeting := "Selam " + name + " :)"
	return greeting
}
