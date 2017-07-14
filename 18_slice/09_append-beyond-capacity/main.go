package main

import "fmt"

func main()  {

	greeting := make([]string, 3, 5)
	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "buenos dias!"

	greeting = append(greeting, "suprabadham")
	greeting = append(greeting, "Zao an")
	greeting = append(greeting, "Ohayou gozaimasu")
	greeting = append(greeting, "gidday")

	fmt.Println(greeting[6])
	fmt.Println(len(greeting))
	fmt.Println(cap(greeting))

}