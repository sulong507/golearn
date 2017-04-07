package work

/*
use chan with no cache to create an goroutine pool
*/

import (
	"sync"
)

type Worker interface {
	Task()
}

type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

func New(maxGoroutine int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}
	p.wg.Add(maxGoroutine)

	for i := 0; i < maxGoroutine; i++ {
		go func() {
			for w := range p.work { //无缓冲的chan,再用了for range后会一直阻塞，直到work通道获取到Worker接口类型的值
				w.Task()
			}
			p.wg.Done()
		}()
	}
	return &p
}

func (p *Pool) Run(w Worker) {
	p.work <- w
}

func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}
