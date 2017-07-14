package main

import "fmt"

func main()  {
	var name string
	fmt.Print("Please type your name: ")
	fmt.Scanln(&name)
	fmt.Println("Hello ", name)
}
