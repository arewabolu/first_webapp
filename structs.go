package main

import "fmt"

// define a new type
// type animal struct {
// 	name  string
// 	class string
// }

type domestic struct {
	animal
	danger string
}

// struct is passed by value

func main() {
	var house domestic
	// iota
	house.animal.name = "dog"
	fmt.Printf("I know the house pets name it %s", house.animal.name)
}
