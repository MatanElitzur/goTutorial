// Package name and folder name must be identical for the go language to recognize the package
package knowledge

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func ConcurrentPrograming1() {
	concurrentPrograming1()
	concurrentPrograming2()
	concurrentProgramingChan1()
	concurrentProgramingChan2()
	concurrentProgramingChan3()
	concurrentProgramingChan4()
	concurrentProgramingChan5()
	concurrentProgramingChan6()
	concurrentProgramingChan7()
	concurrentProgramingMutex1()
	concurrentProgramingMutex2()
	concurrentProgramingMutex3()
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
	wg.Wait() //Main process is blocked with wg.Wait()
}

func concurrentProgramingChan1() {
	c := make(chan int)
	go func() {
		sum := 0
		for i := 0; i < 10; i++ {
			fmt.Println("IDX from go routine func:", i)
			sum += 1
		}
		c <- sum
	}()
	output := <-c //Basiclly this is like wg.Wait() it is blocking the main process (We get the chanel value)
	fmt.Println("Output:", output)
}

type SafeCounter struct {
	mu     sync.Mutex
	NumMap map[string]int
}

func (s *SafeCounter) add(num int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.NumMap["key"] = num
}

// This function will create 100 go routine and will update the number with function add(num int)
// But each go routine will update the value separately cause of the s.mu.Lock()
func concurrentProgramingMutex1() {
	s := SafeCounter{NumMap: make(map[string]int)}
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			s.add(i)
		}(i)
	}
	wg.Wait()
	fmt.Println("concurrentProgramingMutex1:", s.NumMap["key"])
	//Basically we don't know what will be the number in the print cause we don't know which
	//go routine updated the value at the end
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////
func concurrentProgramingMutex2() {
	var tickets int = 500 //Total number of tickets available
	var wg sync.WaitGroup
	mutexPointer := &sync.Mutex{} // Create a pointer to a mutex
	//simulating a lot of users trying to but tickets
	//For simplicity let's assume number of tickets a user can buy is one
	for userId := 0; userId < 2000; userId++ {
		wg.Add(1)
		//Buy ticket for the user with ID userId
		go buyTicket(&wg, mutexPointer, userId, &tickets)
	}
	wg.Wait()
}

/*
a function that simulates a user trying to buy  a ticket.
It sends a request to the manageticket goroutine through ticketChan
*/
func buyTicket(wg *sync.WaitGroup, mutex *sync.Mutex, userId int, remainigTickets *int) {
	defer wg.Done()
	mutex.Lock()
	if *remainigTickets > 0 {
		*remainigTickets-- //User purchases a ticket
		fmt.Printf("User %d purchased a ticket. Tickets remaining: %d \n", userId, *remainigTickets)
	} else {
		fmt.Printf("User %d found no ticket. \n", userId)
	}
	mutex.Unlock()
}

// ///////////////////////////////////////////////////////////////////////////////////////////////////
func concurrentProgramingMutex3() {
	var tickets int = 500 //Total number of tickets available
	var wg sync.WaitGroup
	ticketChan := make(chan int)
	doneChan := make(chan bool)
	go manageTicket(ticketChan, doneChan, &tickets)
	for userId := 0; userId < 2000; userId++ {
		wg.Add(1)
		go buyTicket3(&wg, ticketChan, userId)
	}
	wg.Wait()
	doneChan <- true //There can be a senario that less then 500 users wants to buy the tickets
	//in this case we want to signal the manageTicket go routine that we want to finish with the request
}

/*
A function that is responsiable for managing ticket allocations and responding to user requests.
It listens for incoming requests on a channel (ticketChan) and signals on another channel (doneChan)
when it's time to stop
*/
func manageTicket(ticketChan chan int, doneChan chan bool, tickets *int) {
	for {
		select {
		case user := <-ticketChan: //gets the userId from the channel,
			if *tickets > 0 {
				*tickets--
				fmt.Printf("User %d purchased a ticket. Tickets remaining: %d \n", user, *tickets)
			} else {
				fmt.Printf("User %d found no ticket. \n", user)
			}
		case <-doneChan: //is done triggered
			fmt.Printf("Ticket remaining. \n", *tickets)
		}
	}
}

/*
a function that simulates a user trying to buy  a ticket.
It sends a request to the manageticket goroutine through ticketChan
*/
func buyTicket3(wg *sync.WaitGroup, ticketChan chan int, userId int) {
	defer wg.Done()
	ticketChan <- userId
}

// /////////////////////////////////////////////////////////////////////////////////////////////////////
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
	wg.Wait() //Main process is blocked with wg.Wait()
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

func concurrentPrograming2() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Println("IDX from FIRST func:", i)
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Println("IDX from SECOND func:", i)
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		}
	}()

	wg.Wait() //Main process is blocked with wg.Wait()
	fmt.Println("concurrentPrograming2 ended")
}

// //////////////////////////////////////////////////////////////////////////////////////////
func sayHello(name string, ch chan string) {
	time.Sleep(1 * time.Second) // Simulate some work
	fmt.Println("concurrentProgramingChan2 Hello,", name)
	ch <- "concurrentProgramingChan2 Hello from " + name // Send message to channel
}

func concurrentProgramingChan2() {
	// Create a channel to receive messages
	ch := make(chan string)

	// Launch two goroutines to say hello to different names
	go sayHello("Alice", ch)
	go sayHello("Bob", ch)

	// Receive messages from the channel (order might vary)
	msg1 := <-ch
	msg2 := <-ch

	fmt.Println(msg1)
	fmt.Println(msg2)
}

// ///////////////////////////////////////////////////////////////////////////////
func producer(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i // Send value to the channel
		fmt.Println("concurrentProgramingChan3 Sent:", i)
		time.Sleep(500 * time.Millisecond) // Simulate some work
	}
}

func consumer(ch <-chan int) {
	for i := 0; i < 10; i++ {
		value := <-ch // Receive value from the channel
		fmt.Println("concurrentProgramingChan3 Received:", value)
		time.Sleep(1 * time.Second) // Simulate some work
	}
}

func concurrentProgramingChan3() {
	// Create a buffered channel with a capacity of 10 integers
	ch := make(chan int, 10)

	// Launch a producer goroutine
	go producer(ch)

	// Launch a consumer goroutine
	go consumer(ch)

	// Wait for both goroutines to finish
	time.Sleep(15 * time.Second) // Adjust this time as needed
	fmt.Println("Exiting main function")
}

func concurrentProgramingChan4() {
	orders := make(chan string) //Unbuffered channel

	//Customer placing orders
	//The system can handle only one oreder at a time
	go func() {
		for i := 1; i <= 5; i++ { //For loops simulating 5 customers
			order := fmt.Sprintf("concurrentProgramingChan4 Coffee order #%d", i)
			orders <- order //Blocks until the barista is ready to accept the order
		}
		close(orders) //Once all orders are finish we close the channel
	}()

	//Barista processing orders
	for order := range orders { //The barista waits for a channel
		fmt.Printf("concurrentProgramingChan4 Preparing: %s\n", order)
		time.Sleep(2 * time.Second) //Time taken to prepare the order
		fmt.Printf("concurrentProgramingChan4 Served: %s\n", order)
	} //The for loop brake when the cannel is closed "close(orders)""
}

func concurrentProgramingChan5() {
	orders := make(chan string, 3) //Buffered channel with a capacity of 3

	//Customer placing orders
	//The system can handle up to 3 oreders at a time
	go func() {
		for i := 1; i <= 5; i++ { //For loops simulating 5 customers
			order := fmt.Sprintf("concurrentProgramingChan5 Coffee order #%d", i)
			orders <- order //Only blocks if the buffer is full
		}
		close(orders) //Once all orders are finish we close the channel
	}()

	//Barista processing orders
	for order := range orders { //The barista waits for a channel
		fmt.Printf("concurrentProgramingChan5 Preparing: %s\n", order)
		time.Sleep(2 * time.Second) //Time taken to prepare the order
		fmt.Printf("concurrentProgramingChan5 Served: %s\n", order)
	} //The for loop brake when the cannel is closed "close(orders)""
}

// ////////////////////////////////////////////////////////////////////////////
func doWork(d time.Duration, resch chan string) {
	fmt.Printf("concurrentProgramingChan6 doing work...")
	time.Sleep(d)
	fmt.Printf("concurrentProgramingChan6 work is done!")
	resch <- fmt.Sprintf("concurrentProgramingChan6 the result of the work -> %d", rand.Intn(100))
	wg.Done()
}

var wg *sync.WaitGroup

func concurrentProgramingChan6() {
	start := time.Now()
	resultch := make(chan string)
	wg = &sync.WaitGroup{}
	wg.Add(3)
	go doWork(time.Second*2, resultch)
	go doWork(time.Second*4, resultch)
	go doWork(time.Second*6, resultch)

	go func() {
		for res := range resultch {
			fmt.Println(res)
		}
		fmt.Printf("concurrentProgramingChan6 work took %v seconds\n", time.Since(start))
	}()

	time.Sleep(time.Second)
	wg.Wait()
	close(resultch)

}

func generator(n int, ch chan int) {
	for i := 0; i < n; i++ {
		ch <- i // send a value to the channel
	}
	close(ch) // close the channel after the loop
}

func concurrentProgramingChan7() {
	fmt.Printf("concurrentProgramingChan7 start")
	ch := make(chan int)

	go generator(5, ch)

	for val := range ch {
		fmt.Println(val) // receive values one by one from the channel
	}
	fmt.Printf("concurrentProgramingChan7 ended")

	//panic("concurrentProgramingChan7 panic")
}
