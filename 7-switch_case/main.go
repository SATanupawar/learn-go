package main
 import "fmt"
 import "time"

 func main (){
	day := time.Now().Weekday() // get the current day
	switch day {
	case time.Sunday, time.Saturday: // check if the day is weekend
		fmt.Println("it is weekend and today is", day)
	default: // if the day is not weekend
		fmt.Println("it is weekday and today is", day)
	}

}