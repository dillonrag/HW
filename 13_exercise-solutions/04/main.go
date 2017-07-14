package main

import "fmt"

func main()  {
	var smallNum int
	var largeNum int
	fmt.Println("Please type a small number:")
	fmt.Scan(&smallNum)
	fmt.Println("Please type a large number:")
	fmt.Scan(&largeNum)
	remainder := largeNum % smallNum
	fmt.Println("The remainder is ", remainder)

}