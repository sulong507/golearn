package main

import (
	_ "io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	maxGoroutines   = 25
	pooledResources = 2
)

type dbConnection struct {
	ID int32
}

func (dbConn *dbConnection) Close() error {
	log.Println("Close: connection", dbConn.ID)
	return nil
}

var idCounter int32

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	NewConnection := func() interface{} {
		id := atomic.AddInt32(&idCounter, 1)
		log.Println("Create New Connection", id)
		return &dbConnection{id}
	}

	p := &sync.Pool{New: NewConnection}
	/*
		for i := 0; i < pooledResources; i++ {
			conn := NewConnection
			if conn != nil {
				p.Put(conn)
			}
		}
	*/

	for query := 0; query < maxGoroutines; query++ {
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}
	wg.Wait()
	log.Println("shutdown Program")
}

func performQueries(query int, p *sync.Pool) {
	conn := p.Get()
	if conn == nil {
		log.Println("get connect error.")
		return
	}

	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}
