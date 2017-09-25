package main

import "fmt"
//import "math"

func main()  {
	for a := 1; a < 998/2; a++ {
		for b := 1; b < 998; b++ {
			c := 1000-a-b
			if a*a+b*b == c*c {
				fmt.Println("The triplet is: ", a, b, c)
				fmt.Println("The product is: ", a*b*c)
				break
			}
		}
	}
}