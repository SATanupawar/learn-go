package main

import "fmt"

// if else statement

func main() {
	age := 10

	if age >= 18 {
		fmt.Println("you are an adult")
	} else if age >= 12 {
		fmt.Println("you are a teenager")
	} else {
		fmt.Println("you are a child")
	}

}
