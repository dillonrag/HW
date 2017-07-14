package main

import "fmt"

func greatest(x ...int) int  {
	var g int
	for _, n := range x {
		if n > g {
			g = n
		}
	}
	return g
}

func main()  {
	num := greatest(23, 2, 5, 16)
	fmt.Println(num)
}
