package main

import (
	"log"
	"io"
	"sync/atomic"
	"go-learn/charpter7/pool"
	"sync"
	"time"
	"math/rand"
)

const (
	maxGoRoutines = 25
	pooledResources = 2
)

type dbConnection struct {
	ID int32
}

var idCounter int32

func (dbConn *dbConnection) Close() error {
	log.Println("close connection", dbConn.ID)
	return nil
}

func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("create new connection", id)

	// why not using 'new' ?
	return &dbConnection{id}, nil
}



func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoRoutines)
	
	p, err := pool.New(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
		return
	}

	for query:=0; query<maxGoRoutines; query++ {
		go func(q int) {
//			fmt.Println(q)
			//			p.Acquire()
			performanceQuqery(p, q)
			wg.Done()
		}(query)
	}

	wg.Wait()

	log.Println("shutdown Program....")
	p.Close()
}

func performanceQuqery(p *pool.Pool, q int) {
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}
	defer p.Release(conn)

	time.Sleep(time.Duration(rand.Intn(10)) * time.Microsecond)
	log.Println(q, " is related to connection id ", conn.(*dbConnection).ID)
}
