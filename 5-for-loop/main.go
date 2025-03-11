package main 

import "fmt"

func main() {
	for i := 0; i < 5; i++ {

		if i == 3 {
			continue // skip the iteration
		}
		fmt.Println(i)
	}
}