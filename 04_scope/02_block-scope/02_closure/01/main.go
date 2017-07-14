package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "One ring to rules them all"
		fmt.Println(y)
	}
	//fmt.Println(y)
}