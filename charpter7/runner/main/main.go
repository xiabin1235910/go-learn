package main

import (
	"log"
	"time"
	"go-learn/charpter7/runner"
	"os"
)


const timeout = 3 * time.Second

func main() {
//	log.Print(time.Duration(1))

//	var tt time.Duration = 2;
//	log.Print(tt.Seconds())

	r := runner.New(timeout)
	r.Add(createTask(), createTask(), createTask())
	
	//	r.Start()
	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("Terminating due to timeout.")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("Terminating due to interrupt.")
			os.Exit(2)
		}
	}

	log.Println("Process ended.")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor, Task is #%d", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
