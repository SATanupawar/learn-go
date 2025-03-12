package main

import "fmt"

// basically map is similar to object in javascript

// map is a collection of key-value pairs
func main() {
	
m := make(map[string]int) // map[key-type]value-type / define a map
// add key-value pairs to the map
m["key1"] = 10
m["key2"] = 20
m["key3"] = 30

fmt.Println(m["key1"],m["key2"],m["key3"])	

// delete a key-value pair from the map
// delete(m,"key1")
// fmt.Println(m)

// check if a key is present in the map
// value, ok := m["key1"]
// fmt.Println(value, ok)
}
