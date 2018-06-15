package main

import (
	"os"
	"io/ioutil"
	"go-learn/charpter3/words"
	"fmt"
)

func main() {
	filename := os.Args[1]

	contents, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("There was an error when opening the file", err)
		return
	}

	// variable type convertion from bytes to string
	text := string(contents)
	count := words.CountWords(text)

	fmt.Printf("There are %d words in your text. \n", count)
}
