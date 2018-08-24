package main

import (
	"os"
	"os/signal"
	"fmt"
)

func main() {
	c := make(chan os.Signal, 1)

	fmt.Print("123")
	signal.Notify(c, os.Interrupt)
	fmt.Print("456")
	
	s := <-c
	fmt.Println("Got signal:", s)
}
