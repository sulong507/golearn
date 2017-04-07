package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var (
	counter  int64
	shutdown int64
	wg       sync.WaitGroup
	mutex    sync.Mutex
)

func main() {
	runtime.GOMAXPROCS(1)

	wg.Add(2)

	fmt.Println("Create Gorutines.")
	go incCounter2(1)
	go incCounter2(2)

	//go doWork("A")
	//go doWork("B")

	time.Sleep(1 * time.Second)
	fmt.Println("shutdown now.")
	atomic.StoreInt64(&shutdown, 1)

	fmt.Println("Waiting To Finish...")
	wg.Wait()
	fmt.Println("Terminating Program.")
	fmt.Println("Final Counter: ", counter)
}

//锁住共享资源的方法1：原子函数 AddInt64 LoadInt64 StoreInt64
func incCounter(id int) {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counter, 1)
		runtime.Gosched()
	}
}

//互斥锁
func incCounter2(id int) {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		mutex.Lock()
		{
			value := counter
			runtime.Gosched()
			value++
			counter = value
		}
		mutex.Unlock()
	}
}

func doWork(name string) {
	defer wg.Done()
	for {
		fmt.Printf("Doing %s work\n", name)
		time.Sleep(250 * time.Millisecond)
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shuting %s down\n", name)
			break
		}
	}
}
