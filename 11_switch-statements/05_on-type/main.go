package main

import "fmt"

type Contact struct {
	greeting string
	name     string
}

func SwitchOnType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Contact:
		fmt.Println("Contact")
	default:
		fmt.Println("Unknown")
	}
}

func main()  {
	SwitchOnType(7)
	SwitchOnType("Raghunandan")
}
