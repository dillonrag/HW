package main

import "fmt"

func main() {

	myFriendsName := "Al"

	switch {
	case len(myFriendsName) == 2:
		fmt.Println("Sup friend with name of lenght 2")
	case myFriendsName == "Tim":
		fmt.Println("Sup Tim")
	case myFriendsName == "Jenny":
		fmt.Println("Sup Jenny")
	case myFriendsName == "Marcus", myFriendsName == "Medhi":
		fmt.Println("Your name is either Marcus or Medhi")
	case myFriendsName == "Julian":
		fmt.Println("Sup Julian")
	case myFriendsName == "Sushant":
		fmt.Println("Sup Sushant")
	default:
		fmt.Println("nothing matched; this is the default")
	}
}
