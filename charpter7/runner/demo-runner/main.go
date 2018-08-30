package main

import (
	"time"
	"go-learn/charpter7/runner/demo-runner/runner"
	"log"
	"os"
)

func main() {
	var r *runner.Runner = runner.New(3 * time.Second)
	r.Add(createTask, createTask,createTask)

	if err := r.Start(); err != nil {
		switch(err) {
		case runner.ErrTimeout:
			log.Println("terminate due to timeout")
			os.Exit(2)
		case runner.ErrInterrupt:
			log.Println("terminate due to interrupt")
			os.Exit(1)
		}
	}
	
}

func createTask(id int) {
	log.Println("task id is", id)
	time.Sleep(time.Duration(id) * time.Second)
}
