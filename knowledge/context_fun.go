package knowledge

import (
	"context"
	"fmt"
	"log"
	"time"
)

func enrichContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "request-id", "12345")
}

func doSomethingCool(ctx context.Context) {
	rID := ctx.Value("request-id")
	fmt.Println(rID)
	for {
		select {
		case <-ctx.Done(): //Done is called when the timeout is trigger
			fmt.Println("timed out")
			return
		default:
			fmt.Println("doing something cool")
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func ExecuteContext_1() {
	fmt.Println("Go context Tutorial")
	//ctx1 := context.Background() //creation of context
	//ctx, cancel := context.WithTimeout(ctx1, 2*time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second) //After 2 secounds the ctx.Done() is trigger,, casue the time has ended
	defer cancel()                                                          // We want to cancel the context at the end, if we invoke cancel() without defer it will cancel right away
	fmt.Println(ctx.Err())                                                  //Will print <nil> --> Because we called it before we called ctx.Done()
	ctx = enrichContext(ctx)
	go doSomethingCool(ctx)
	select {
	case <-ctx.Done(): //Done is called when the timeout is trigger
		fmt.Println("oh no, I'v exceeded the deadline")
		fmt.Println(ctx.Err()) //Will print: context deadline exceeded --> becasue we called it after ctx.Done()
	}
	time.Sleep(2 * time.Second)
}

// ////////////////////////////////////////////////////////////
const shortDuration = 1 * time.Millisecond //At the select will print ExecuteContext_11 ctx.Done()
//const shortDuration = 3 * time.Millisecond //At the select will print overslept

func ExecuteContext_11() {
	//Pass a context with a timeout to tell a blocking function that is
	//shouled abandon its work after the timeout elapses
	ctx, cancel := context.WithTimeout(context.Background(), shortDuration)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println("ExecuteContext_11 ctx.Done()")
		fmt.Println(ctx.Err()) //prints "context deadline exceeded"
	}
}

func ExecuteContext_12() {
	ch := make(chan struct{}) // creates an unbuffered channel, to send only signals, Not to pass data
	//ch := make(chan string) (channel for strings)
	//ch := make(chan int, 10) (buffered channel, channel for integers with a capacity of 10)
	//sendCh := make(chan<- string) (Directional Channels:, send-only channel for strings)
	run := func(ctx context.Context) {
		n := 1
		for {
			select {
			case <-ctx.Done(): //ctx is cancelled, we close ch
				fmt.Println("ExecuteContext_12 existing")
				close(ch)
				return // returning no to leak the goroutine
			default:
				time.Sleep(time.Millisecond * 300)
				fmt.Println("ExecuteContext_12", n)
				n++
			}
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("ExecuteContext_12", "goodbye")
		cancel() // cancels ctx
	}()

	go run(ctx)
	fmt.Println("ExecuteContext_12 waiting to cancel ...")
	<-ch
	fmt.Println("ExecuteContext_12 bye")
}

type jwt string

const auth jwt = "JWT"

func ExecuteContext_13() {
	ctx := context.WithValue(context.Background(), auth, "Bearer hi")
	bearer := ctx.Value(auth) // get the actual value from the key==auth
	str, ok := bearer.(string)
	if !ok {
		log.Fatalln("ExecuteContext_13 not a string")
	}
	fmt.Println("ExecuteContext_13 value:", str)
}

// I want to create a context with a value, and then I want to retrieve that value from the context
