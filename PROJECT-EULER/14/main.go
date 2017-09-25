package main

import "fmt"

func main()  {
	lens := []int{}
	for a := 1; a < 1000000; a++ {
		lens = append(lens, Collatz(a))
	}
	max := 0
	for _, v := range lens {
		if v > max {
			max = v
		}
	}
	for i := 0; i <= len(lens); i++ {
		if lens[i] == max {
			fmt.Println("The number with the longest Collatz chain is: ", i+1)
			break
		}
	}

}

func Collatz(start int) int{
	xs := []int{}
	xs = append(xs, start)
	for n := start; n != 1; {
		if n % 2 == 0 {
			n = n/2
			xs = append(xs, n)
		} else {
			n = 3*n+1
			xs = append(xs, n)
		}
	}
	return len(xs)
}
