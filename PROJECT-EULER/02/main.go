package main

import "fmt"

func main()  {
	xs := []int{1,2}
	i := 2
	fibSum := 0
	evenSum := 0
		for fibSum < 4000000 {
			fibSum = xs[i-2]+xs[i-1]
			xs= append(xs, fibSum)
			if fibSum < 4000000 && fibSum % 2 ==0 {
				evenSum+= fibSum
			}
			i++
		}
	fmt.Println("The sum is: ", evenSum)
}
