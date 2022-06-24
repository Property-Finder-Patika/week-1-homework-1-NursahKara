package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Args[1]            // arguments are defined
	greeting := createGreet(name) // function assigned to variable greeting
	fmt.Printf("%s\n", greeting)

}

func createGreet(name string) string { // Function defined with return type string
	return "Selam " + name + " :)"
}
