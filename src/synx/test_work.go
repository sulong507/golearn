package main

import (
	"log"
	"sync"
	"synx/work"
	"time"
)

var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

type namePrinter struct {
	name string
}

func (m *namePrinter) Task() {
	log.Println(m.name)
	time.Sleep(100 * time.Millisecond)
}

func init() {
	log.SetPrefix("WORKER:")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	p := work.New(5) //启动5个goroutine，每个都阻塞在从work通道中获取worker，当执行Run方法的时候，goroutine会从work通道中获取到值，才开始执行。
	var wg sync.WaitGroup
	wg.Add(100 * len(names))

	defer func() {
		log.Println("in defer")
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	handle()

	for i := 0; i < 100; i++ {
		for _, name := range names {
			np := namePrinter{
				name: name,
			}
			go func() {
				//***这里必须注意：虽然Run的参数是Worker接口，但是在调用任务的时候，是通过np.Task()调用的，所以必须传指针(&np.Task())。
				p.Run(&np)
				wg.Done()
			}()
		}
	}
	wg.Wait()
	p.Shutdown()
}

func handle() {
	log.Println("in handle")
	panic("has a exception")
	log.Println("after handle")
}
