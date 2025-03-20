package main 
import (
	"fmt"
	"time"
)

// func proccesNum(numChan chan int){
// 	for num := range numChan {
// 		fmt.Println("received number" , num)
// 		time.Sleep(time.Second * 1)
// 	}

// }

// func sum (result chan int , num1 int , num2 int){
// 	numResult := num1 + num2
// 	result <- numResult
// }	

// func task (done chan bool){
// 	defer func (){ done <- true}()
// 	fmt.Println("task is done")
// }

func sendEmail (emailChan chan string , done chan bool){
	defer func (){ done <- true}()
	 for email := range emailChan {
		fmt.Println("sending email to" , email)
		time.Sleep(time.Second * 1)
	}

	
}

func main (){
	fmt.Println("Hello, World!")

// messageChan := make(chan string)

// messageChan <- "Hello, World!"  // send and channels are blocking

// msg := <-messageChan

// fmt.Println(msg)


// numChan := make(chan int)

// go proccesNum(numChan)

//  for {
// 	numChan <- rand.Intn(100)
//  }

// result := make (chan int)

// go sum(result , 1 , 2)

// fmt.Println(<-result)

// done := make(chan bool)

// go task(done)

// fmt.Println(<-done)

	// sendEmail(emailChan , done)


emailChan := make(chan string , 100)

// emailChan <- "satyampawar91@gmail.com"

done := make(chan bool)

go sendEmail(emailChan , done)

// fmt.Println(<-emailChan)

for i := 0 ; i <5 ; i++ {
	emailChan <- fmt.Sprintf("user%d@example.com" , i)
}
fmt.Println("all emails sent")
// close(emailChan) this is important to close the channel
close(emailChan)

<-done

}
