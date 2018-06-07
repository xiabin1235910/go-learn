package matcher

import (
	"go-learn/charpter2/search"
	"errors"
	"net/http"
	"encoding/xml"
	"fmt"
)

type (
	// todo 
	rssDocument struct {}
)

type rssMatcher struct{}

func (m rssMatcher) Retrieve(feed *search.Feed) (*rssDocument, error) {
	if feed.URI == "" {
		return nil, errors.New("URI cannot be empty")
	}

	res, err := http.Get(feed.URI)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Http Response Error %d\n", res.StatusCode)
	}
	
	defer res.Body.Close()
	
	var document rssDocument

	err = xml.NewDecoder(res.Body).Decode(&document)
	return &document, err
}
