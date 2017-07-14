package main

import "fmt"

func main() {
	a := 43
	fmt.Println(a) //43
	fmt.Println(&a)

	var b *int = &a

	fmt.Println(b)
	fmt.Println(*b) //43

	*b = 42 //b says, "The value at this address, change it to 42"
	fmt.Println(a) //42

}

//this is useful
//we can pass a memory address instead of a bunch of values (we can pass a reference
//and then we can still change the value of whatever is stored at the meme address
// this makes our programs more performant
// we don't have to pass around large amounts of data
//we only have to pass around addresses

//everything is PASS BY VALUE in go
//when we pass a mem addy, we are passing a value