package main

import (
	"log"
	"time"
)

func main() {
	log.Print(time.Duration(1))

	var tt time.Duration = 2;
	log.Print(tt.Seconds())
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor, Task is #%d", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
