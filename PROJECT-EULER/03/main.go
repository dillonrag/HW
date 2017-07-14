package main

import "fmt"

func isPrime(n int) bool  {
	var p bool = false
	for i := 2; i*i <= n; i+=2{
		if n%i != 0{
			p = false
			return p
		} else{
			p = true
		}
	}
	return p
}

func main()  {
	oddFactors := []int{}
	primes := []int{1}
	number := 13195
	for i := 1; i < number/2; i+= 2 {
		if number % i == 0 {
			oddFactors = append(oddFactors, i)
		}
	}
	fmt.Println(oddFactors)
	for _, j:= range oddFactors {
		if isPrime(j){
			primes = append(primes, j)
		}
	}
	fmt.Println(primes)
}
