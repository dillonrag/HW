package main

import "fmt"

func main() {
	switch "Jenny" {
	case "Daniel":
		fmt.Println("Wassup Daniel")
	case "Medhi":
		fmt.Println("Wassup Medhi")
	case "Jenny":
		fmt.Println("Wassup Jenny")
		fallthrough
	case "Jared":
		fmt.Println("Wassup Jared")
		fallthrough
	case "Amy":
		fmt.Println("Wassup Amy")
	case "Charles":
		fmt.Println("Wassup Charles")
	}
}