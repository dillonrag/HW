package main

import "fmt"
import "math"

func main()  {
	a,b := 1,1
	sum := 12

	for {
		c := math.Sqrt(float64(a * a + b * b))
		for a+b+int(c) <= sum {
		for a+b+int(c) <= sum {
			if a + b + int(c) == sum {
				fmt.Println("The triplet is", a, b, c)
				break
			}
		}
			b++
	}
		a++
	}
}

//UNFINISHED