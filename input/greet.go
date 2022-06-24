package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Println("Please enter your name: ")
	fmt.Scanln(&name)
	greeting := createGreet(name) // function assigned to variable greeting
	fmt.Printf("%s\n", greeting)
}

func createGreet(name string) string { // Function defined with return type string
	return "Selam " + name + " :)"
}
