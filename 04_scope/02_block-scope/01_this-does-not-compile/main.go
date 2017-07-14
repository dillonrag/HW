package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	foo()
}

func foo() {
	//no access to x, so it doesn't compile
	fmt.Println(x)
}