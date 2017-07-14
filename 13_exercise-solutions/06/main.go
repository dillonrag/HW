package main

import "fmt"

func main()  {
	for i := 1; i < 100; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				fmt.Println(i, " --FizzBuzz")
			} else {
				fmt.Println(i, " --Fizz")
			}
		} else if i%5 == 0 {
			fmt.Println(i, " --Buzz")
		} else {
			fmt.Println(i)
		}

	}
}
