package main

import "fmt"

func main() {
	fmt.Println("Selam")
	greet("Nurşah")

}

func greet(name string) { // Function defined with no return
	fmt.Printf(name)
}
