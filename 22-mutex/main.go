package main 

import (
	"fmt"
	
)

type view struct {
	view int
}

func ( p *view ) inc(){
	p.view++
}


func main (){
	fmt.Println("Hello, World!")

	myPost := view{view: 0}

	for i := 0 ; i < 100 ; i++ {
		 myPost.inc()
	}
	 
	

	fmt.Println(myPost.view)

}