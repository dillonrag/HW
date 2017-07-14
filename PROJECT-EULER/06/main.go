package main

import "fmt"

func main()  {
	susq := 0
	sum := 0

	for i := 0; i <= 100; i++ {
		susq += i*i
		sum += i
	}
	sqsu := sum*sum

	fmt.Println("Sum of Squares =", susq)
	fmt.Println("Square of Sum =", sqsu)
	fmt.Println("Difference =", sqsu-susq)
}
