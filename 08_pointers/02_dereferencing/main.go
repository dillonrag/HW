package main

import "fmt"

func main() {
	a := 43
	fmt.Println(a) //43
	fmt.Println(&a)

	var b *int = &a

	fmt.Println(b)
	fmt.Println(*b) //43
}

//b is an int pointer
//b points to the memory address where an int is stored
//to see the value in that memory address, add an * in front of b
//this is known as dereferencing
//the * is an operator in this case