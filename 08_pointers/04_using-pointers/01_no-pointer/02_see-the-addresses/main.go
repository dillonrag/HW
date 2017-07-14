package main

import "fmt"

func zero(z int) {
	fmt.Printf("%p\n", &z) // address in func zero
	fmt.Println(&z) //addres in fuc zero
	z = 0
}

func main() {
	x := 5
	fmt.Printf("%p\n", &x) // address in main
	fmt.Println(&x) //addres in main
	zero(x)
	fmt.Println(x)
}