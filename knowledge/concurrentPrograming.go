// Package name and folder name must be identical for the go language to recognize the package
package knowledge

import (
	"fmt"
	"sync"
	"time"
)

func ConcurrentPrograming1() {
	concurrentPrograming1()
	concurrentProgramingSelect()
	concurrentProgramingLooping()
}

// type WaitGroup
// func (wg WaitGroup) Add(delta int) //increment counter by delta
// func (wg WaitGroup) Done() //decrement counter by 1
// func (wg WaitGroup) Wait() //wait till counter is zero

func concurrentPrograming1() {
	//Concurrency
	// Create a waitGroup variable to tell the main run to wait until the concurrent task is completed
	var wg sync.WaitGroup
	//With channels we can pass info between go routines
	//make a channels that works with integer
	ch := make(chan int)
	//ch := make(chan string) //Example of creating a channel of type string, so we can pass string between the go routines
	//We call Add with 1 argument indicate that we have 1 task to wait upon
	wg.Add(1)
	//Create an anonymous func go routine with the add key word "go"
	//Tells the go scheduler that I want to run this go routine in the future
	go func() {
		fmt.Println("do some async thing #1")
		ch <- 42 //Send a message into the channel
		//At the end of the go routine we call Done() to indicate that we finish with the the routine
		//wg.Done()
	}() //The () basically invoke the function immediately

	go func() { // Another go routine
		fmt.Println("do some async thing #2")
		fmt.Println(<-ch) //receive the message back from the channel
		wg.Done()
	}() // The scheduler knows that there is a channel between them and therefor there is only 1 --> wg.Add(1)

	// After that we invoked the go routine, we call Wait() in the main run
	//It tells the runtime to wait for the go routine invocation
	//Meanwhile the sceduler receive the request to invoke the go routine
	//So when wg.Done() will be call and decrement by 1 so now the index is on zero
	//And now the wait block will stop to block and the code will continue in the main run.
	wg.Wait()
}

func concurrentProgramingSelect() {
	ch1, ch2 := make(chan string), make(chan string)

	// go func() {
	// 	fmt.Println("do some async thing #1")
	// 	ch1 <- "message to channel 1" //Send a message into the channel
	// 	//At the end of the go routine we call Done() to indicate that we finish with the the routine
	// 	//wg.Done()
	// }() //The () basically invoke the function immediately

	go func() {
		fmt.Println("do some async thing #2")
		ch2 <- "message to channel 2" //Send a message into the channel
		//At the end of the go routine we call Done() to indicate that we finish with the the routine
		//wg.Done()
	}() //The () basically invoke the function immediately

	//Stop the time for 10 Millisecond the scheduler will understand that he have 2 gos routine to invoke
	//Then the go routine will invoke but after 10 Millisecond the main runtime will come back
	//So we are not sure if both routine run or which one of them run
	time.Sleep(10 * time.Millisecond)

	select {
	case msg := <-ch1: //first case will receive a message from channel 1
		fmt.Println(msg)
	case msg := <-ch2: //secound case will receive a message from channel 2
		fmt.Println(msg)
	default:
		fmt.Println("no messages available") // if we don't add a default and we will not have the routines we will get a deadlock
	}

}

func concurrentProgramingLooping() {
	ch := make(chan int)

	go func() {
		fmt.Println("knowledge-->concurrentProgramingLooping-->concurrentProgramingLooping: do some async thing #1")
		for i := 0; i < 10; i++ {
			ch <- i //Send a message into the channel
		}
		close(ch) //Closes a channel, no more messages can be sent, then the loop in line 99 will close cause no more message can be received
	}() //The () basically invoke the function immediately

	for msg := range ch {
		//Every iteration of the channel loop, it will pull one message from the channel
		// The msg variable will get the message from the ch (channel)
		fmt.Printf("knowledge-->errorManagment-->divideNumber3 %v\n", msg)
	}

}
