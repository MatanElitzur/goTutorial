package knowledge

import (
	"context"
	"fmt"
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

func ExecuteContext() {
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
