package main

import "fmt"

func main()  {
	mult := 1
	for {
		k := 0
		for j := 1; j <= 20; j++ {
			if mult % j == 0 {
				k++
			} else {
				mult++
			}
		}
		if k == 20 {
			fmt.Println(mult)
			break
		}
	}
}
