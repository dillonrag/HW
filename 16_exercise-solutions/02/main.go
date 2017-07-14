package main

import "fmt"

func main()  {
	half := func(x int) (int, bool) {
		return x/2, x%2 ==0
	}
	h, even := half(2)
	fmt.Println(h, even)
}