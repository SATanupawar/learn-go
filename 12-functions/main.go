package main

import "fmt"

// create a  Add function
func add (a int , b int )int{
	return a + b
}

func main() {

	result := add(8,4) // call the add function

	fmt.Println(result) // print the result
}
