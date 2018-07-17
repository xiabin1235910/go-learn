package main

import (
	"os"
	"os/signal"
	"fmt"
)

func main() {
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	s := <-c
	fmt.Println("Got signal:", s)
}
