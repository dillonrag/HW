package main

import "fmt"

func main() {
	switch "Jenny" {
	case "Tim", "Jenny":
		fmt.Println("Wassup Time, or ,err, Jenny")
	case "Marcus", "Medhi":
		fmt.Println("Both of your names start with an M")
	case "Julian", "Sushant":
		fmt.Println("Wassup Julian/Sushant")
	}
}
