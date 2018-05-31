package search

import (
	"log"
)

var matchers = make(map[string]Matcher)

func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher has already exists")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
