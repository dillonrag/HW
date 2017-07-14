package main

import "fmt"

func main()  {
	if !false {
		fmt.Println("This ran")
	}

	if !true {
		fmt.Println("This did not run")
	}
}
