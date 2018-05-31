package search

import (
	"log"
)

type Result struct {
	Field string
	Content string
}

type Matcher interface {
	// not very clear about the parameter of function Search
	// TODO
	Search() ([]*Result, error)
}

func Match(matcher Matcher, results chan *Result) {
	searchResults, err := matcher.Search()

	if err != nil {
		log.Println(err)
		return
	}

	// write result to channel
	for _, result := range searchResults {
		results <- result
	
	}
}
