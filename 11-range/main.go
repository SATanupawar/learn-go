package main

import "fmt"

// range is used to iterate over the slice, map, string, etc.

func main() {
//-----------slices-----------

// 	nums := []int{1,2,3,4,5} // define a slice
// 	nums = append(nums, 6) // add an element to the slice
// 	fmt.Println(nums)
// for i := 0; i < len(nums); i++ { // for loop to iterate over the slice
// 	fmt.Println(nums[i])
// 	}

// 	for i ,num := range nums { // range is used to iterate over the slice /  i is the index and num is the value
// 		fmt.Println(i, num)
// 	}

	//----------maps----------

	// m := map[string]int{"key1": 1, "key2": 2, "key3": 3} // define a map
	// fmt.Println(m)

	// for k ,v := range(m) {
	// 	fmt.Println(k, v)
	// }

//-----------strings-----------

	for i , c := range("hello") {
		fmt.Println(i, string(c))
	}

}	