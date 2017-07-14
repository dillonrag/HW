package main

import "fmt"

func main()  {
	m := make([]string, 1, 25)
	fmt.Println(m) // []
	changeMe(m)
	fmt.Println(m) //[Dillon]
}

func changeMe(z []string)  {
	z[0] = "Dillon"
	fmt.Println(z) //[Dillon]
}