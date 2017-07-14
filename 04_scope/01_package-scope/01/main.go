package main

import "fmt"

var x int = 42
//'x' has package level scope

func main() {
	fmt.Println(x)
	foo()
}

func foo() {
	fmt.Println(x)
}