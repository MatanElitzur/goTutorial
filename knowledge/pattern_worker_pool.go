package knowledge

import (
	"fmt"
	"sync"
	"time"
)

// Task definition
//
//	type Task struct {
//		ID int
//	}
type Task interface {
	Process()
}

// Email task definition
type EmailTask struct {
	Email       string
	Subject     string
	MessageBody string
}

// Way to process the tasks
//
//	func (t *Task) Process() {
//		fmt.Printf("WorkerPool_01 Processing task %d\n", t.ID)
//		//simulate a time consuming process
//		time.Sleep(2 * time.Second)
//	}
func (t *EmailTask) Process() {
	fmt.Printf("WorkerPool_01 Sending email to %d\n", t.Email)
	//simulate a time consuming process
	time.Sleep(2 * time.Second)
}

// Image processing task
type ImageProcessingTask struct {
	ImageUrl string
}

func (t *ImageProcessingTask) Process() {
	fmt.Printf("WorkerPool_01 Processing the image %d\n", t.ImageUrl)
	//simulate a time consuming process
	time.Sleep(5 * time.Second)
}

// worker pool definition
type WorkerPool struct {
	Tasks       []Task
	concurrency int
	tasksChan   chan Task
	wg          sync.WaitGroup
}

// Functions to execute the worker pool
func (wp *WorkerPool) worker() {
	for task := range wp.tasksChan { // Eatch interation we take a task
		task.Process()
		wp.wg.Done()
	}
}

func (wp *WorkerPool) Run() {
	//Initialize the tasks channel
	wp.tasksChan = make(chan Task, len(wp.Tasks))

	//Start workers, the number of workers that we start is i < wp.concurrency
	for i := 0; i < wp.concurrency; i++ {
		go wp.worker()
	}
	//Send tasks to the tasks channel
	wp.wg.Add(len(wp.Tasks)) //Matan - Add all the tasks to the waitgroup
	for _, task := range wp.Tasks {
		wp.tasksChan <- task //Matan - add all the tasks to the chanel of type Task
	}
	close(wp.tasksChan) //Signal that no more tasks will be sent

	//Wait for all tasks to finish
	wp.wg.Wait()
}

func WorkerPool_01() {
	//Create new tasks
	// tasks := make([]Task, 20)
	// for i := 0; i < 20; i++ {
	// 	tasks[i] = Task{ID: i + 1}
	// }
	tasks := []Task{
		&EmailTask{Email: "email1@codeheim.io", Subject: "test", MessageBody: "test"},
		&ImageProcessingTask{ImageUrl: "/images/sample1.jpg"},
		&EmailTask{Email: "email2@codeheim.io", Subject: "test", MessageBody: "test"},
		&ImageProcessingTask{ImageUrl: "/images/sample2.jpg"},
		&EmailTask{Email: "email3@codeheim.io", Subject: "test", MessageBody: "test"},
		&ImageProcessingTask{ImageUrl: "/images/sample3.jpg"},
		&EmailTask{Email: "email4@codeheim.io", Subject: "test", MessageBody: "test"},
		&ImageProcessingTask{ImageUrl: "/images/sample4.jpg"},
		&EmailTask{Email: "email5@codeheim.io", Subject: "test", MessageBody: "test"},
		&ImageProcessingTask{ImageUrl: "/images/sample5.jpg"},
	}
	//Create a worker pool
	wp := WorkerPool{
		Tasks:       tasks,
		concurrency: 5, //Number of workers that can run at a time
	}

	//run the pool
	wp.Run()
	fmt.Println("WorkerPool_01 All tasks have been processed")
}
