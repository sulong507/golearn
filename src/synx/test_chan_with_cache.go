package main

import (
	"fmt"
	"image"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4
	taskLoad         = 10
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	task := make(chan string, taskLoad)

	wg.Add(numberGoroutines)

	//四个goroutine,四个worker
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(task, gr)
	}

	//往chan中放入10个task
	for post := 1; post <= taskLoad; post++ {
		task <- fmt.Sprintf("Task: %d", post)
	}

	//任务分配给worker后，关闭chan（此时任务应该还没有执行完成）
	close(task)
	fmt.Println("task chan shutdown now")

	wg.Wait()

	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}

func worker(tasks chan string, worker int) {
	defer wg.Done()

	for {
		task, ok := <-tasks
		if !ok {
			fmt.Printf("Worker: %d : shutting Down\n", worker)
			return
		}

		fmt.Printf("Worker: %d : Started %s\n", worker, task)

		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		fmt.Printf("Worker: %d : Completed %s\n", worker, task)
	}
}
