package main

import (
	"fmt"
	"unicode/utf8"
	"strconv"
)

func main()  {
	xs := []int{1,1}
	i := 2
	fib := 0
	for {
		fib = xs[i-2]+xs[i-1]
		xs = append(xs, fib)
		i++
		if utf8.RuneCountInString(strconv.Itoa(fib)) == 10{ //should come up with solution if you change value to 1000. int cannot store a number that big though....
			fmt.Println(xs)
			fmt.Println(len(xs))
			break
		}
	}
	/*fmt.Println(fib)
	digits := utf8.RuneCountInString(strconv.Itoa(fib))
	fmt.Println(digits)*/

}

//UNFINISHED