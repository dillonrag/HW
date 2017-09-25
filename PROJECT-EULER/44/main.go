package main

import (
	"fmt"
	"reflect"
)

func main()  {
	//Create slice with Pentagonal values
	xs := []int{}
	for n := 1; n <= 100; n++ {
		pn := n*(3*n-1)/2
		xs = append(xs, pn)
	}
	//fmt.Println(xs)

	sums := sumCheck(xs)
	diffs := difCheck(xs)

	matches := [][]int{}
	for _, v := range sums {
		for _, w := range diffs {
			if reflect.DeepEqual(v, w) {
				matches = append(matches, w)
			}
		}
	}
	fmt.Println(matches)
	fmt.Println(diffs)
	fmt.Println(sums)
}

func sumCheck(s []int) [][]int {
	pairs := [][]int{}
	for i := 0; i < len(s); i++ {
		for j := i+1; j < len(s)-1; j++ {
			k := s[i]+s[j]
			for _, v := range s {
				if v == k {
					pair := []int{i+1,j+1}
					pairs = append(pairs, pair)
				}
			}
		}
	}
	return pairs
}

func difCheck(s []int) [][]int {
	pairs := [][]int{}
	for i := 0; i < len(s); i++ {
		for j := i+1; j < len(s)-1; j++ {
			k := s[j]-s[i]
			for _, v := range s {
				if v == k {
					pair := []int{i+1,j+1}
					pairs = append(pairs, pair)
				}
			}
		}
	}
	return pairs
}

//UNFINSHED