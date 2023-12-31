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
	playWithPets()
	concurrentProgramingUsingBufferedChannels()
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
		wg.Done()         // //At the end of the go routine we call Done() to indicate that we finish with the the routine
	}() // The scheduler knows that there is a channel between them and therefor there is only 1 --> wg.Add(1)

	// After that we invoked the go routine, we call Wait() in the main run
	//It tells the runtime to wait for the go routine invocation
	//Meanwhile the sceduler receive the request to invoke the go routine
	//So when wg.Done() will be call and decrement by 1 so now the index is on zero
	//And now the wait block will stop to block and the code will continue in the main run.
	wg.Wait()
}

func concurrentProgramingSelect() {
	ch1, ch2 := make(chan string), make(chan string) //Creating two chanels that work with type string

	// go func() {
	// 	fmt.Println("do some async thing #1")
	// 	ch1 <- "message to channel 1" //Send a message into the channel
	// 	//At the end of the go routine we call Done() to indicate that we finish with the the routine
	// 	//wg.Done()
	// }() //The () basically invoke the function immediately

	go func(str string) {
		fmt.Printf("do some async thing #2 message:%s\n", str)
		ch2 <- "message to channel 2" //Send a message into the channel
		//At the end of the go routine we call Done() to indicate that we finish with the the routine
		//wg.Done()
	}("RIGHT NOW!") //The () basically invoke the function immediately

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
	ch := make(chan int) //Creating a chanel that work with type integer

	go func() {
		fmt.Println("knowledge-->concurrentProgramingLooping-->concurrentProgramingLooping: do some async thing #1")
		for i := 0; i < 10; i++ {
			fmt.Printf("knowledge-->concurrentProgramingLooping-->send message %v\n", i)
			ch <- i //Send a message into the channel
		}
		close(ch) //Closes a channel, no more messages can be sent, then the loop in line 99 will close cause no more message can be received
	}() //The () basically invoke the function immediately

	for msg := range ch {
		//Every iteration of the channel loop, it will pull one message from the channel
		// The msg variable will get the message from the ch (channel)
		fmt.Printf("knowledge-->concurrentProgramingLooping receive message: %v\n", msg)
	}
}

// //ch := make(chan int, 5) //create a buffered channel that has internal capacity of 5 and work with integers
// Example how to create a buffered channel, that has a capacity and that can store 5 messages in the cannnel with out using a reciver.
func concurrentProgramingUsingBufferedChannels() {
	wg := &sync.WaitGroup{} // Creating a pointer to waitgroup
	ch := make(chan int, 1) // Create a buffered channel, that can take 2 messages
	wg.Add(2)               //Takes 2 go routines

	go func(ch <-chan int, wg *sync.WaitGroup) { //ch <-chan int -->This is only a go routine to receive a channel message
		fmt.Println("concurrentPrograming-->concurrentProgramingUsingBufferedChannels #1")
		if msg, ok := <-ch; ok {
			fmt.Println(msg, ok) //ok as false indicate that the chanel is closed
			//fmt.Println(<-ch) //Get the info from the channel and print it.
		}
		wg.Done() //At the end of the go routine we call Done() to indicate that we finish with the the routine
	}(ch, wg)

	go func(ch chan<- int, wg *sync.WaitGroup) { //ch chan<- int -->This is only a go routine to send a channel message
		fmt.Println("concurrentPrograming-->concurrentProgramingUsingBufferedChannels #2")
		//With buffered channel we can send more then one message into a channel
		ch <- 432 //Write to the channel
		//time.Sleep(5 * time.Millisecond)
		//fmt.Println(<-ch) //Get the info from the channel and print it. (I can do it if the routine will be: "go func(ch chan int, wg *sync.WaitGroup)", look for 'CHANNEL TYPES' in the code)
		ch <- 234 //Write to the channel
		wg.Done() //At the end of the go routine we call Done() to indicate that we finish with the the routine
	}(ch, wg)
	// After that we invoked the go routine, we call Wait() in the main run
	//It tells the runtime to wait for the go routine invocation
	//Meanwhile the sceduler receive the request to invoke the go routine
	//So when wg.Done() will be call and decrement by 1 so now the index is on zero
	//And now the wait block will stop to block and the code will continue in the main run.
	wg.Wait()
}

////////////// CHANNEL TYPES /////////////////////
//func myFunction(ch chan int){} //bidirectional channel (send and get messages)
//func myFunction(ch chan<- int){} //Only send channel message
//func myFunction(ch <-chan int){} //Only receive channel message

// ////////////Mutexes --> How to handle shared memory/////////////
// Mutex --> A mutual exclusion lock
// This solve the issue when you read and write to the same memory
var pets []string = []string{"cat", "dog", "lion"}

func playWithPets() {
	fmt.Println("start concurrentPrograming-->playWithPets()")
	mutexPointer := &sync.Mutex{} // Create a pointer to a mutex
	//mutexPointer := &sync.RWMutex{} // Create a RWMutex mutex pointer, allow multipale readers to read and lock the write

	go func(m *sync.Mutex) {
		//m.RLock() // When using RWMutex will allow multipale readers to read
		m.Lock()
		for _, v := range pets { // i == index, v == value
			fmt.Println(v)
		}
		m.Unlock()
		//m.RUnlock() //When using RWMutex
	}(mutexPointer)

	go func(m *sync.Mutex) {
		m.Lock()
		pets = append(pets, "donkey") // add elements to the slice
		fmt.Println(pets)
		pets = append(pets, "horse") // add elements to the slice
		fmt.Println(pets)
		pets = append(pets, "hippo") // add elements to the slice
		fmt.Println(pets)
		m.Unlock()
	}(mutexPointer)

	time.Sleep(10 * time.Millisecond)
	fmt.Println("end concurrentPrograming-->playWithPets()")
}
