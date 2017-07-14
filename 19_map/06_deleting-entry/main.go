package main

import "fmt"

func main()  {

	myGreeting := map[int]string{
		0: "Good morning!",
		1: "Bonjour!",
		2: "Buenos Dias!",
		3: "Bongiorno!",
	}

	fmt.Println(myGreeting)
	delete(myGreeting, 2)
	fmt.Println(myGreeting)

}
