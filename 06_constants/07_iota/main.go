package main

import "fmt"

const (
	_ = iota //0
	KB = 1 << (iota * 10) // shifts byte to the left
	MB = 1 << (iota * 10) // shifts byte to the left
	GB = 1 << (iota * 10) // shifts byte to the left
)

func main() {
	fmt.Println("binary\t\tdecimal")
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)

}
