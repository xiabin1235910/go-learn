package matcher

import (
	"go-learn/charpter2/search"
	"errors"
	"net/http"
	"encoding/xml"
	"fmt"
	"regexp"
)

type (
	// item defines the fields associated with the item tag
	// in the rss document.
	item struct {
		XMLName     xml.Name `xml:"item"`
		PubDate     string   `xml:"pubDate"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
		Link        string   `xml:"link"`
		GUID        string   `xml:"guid"`
		GeoRssPoint string   `xml:"georss:point"`
	}

	// image defines the fields associated with the image tag
	// in the rss document.
	image struct {
		XMLName xml.Name `xml:"image"`
		URL     string   `xml:"url"`
		Title   string   `xml:"title"`
		Link    string   `xml:"link"`
	}

	// channel defines the fields associated with the channel tag
	// in the rss document.
	channel struct {
		XMLName        xml.Name `xml:"channel"`
		Title          string   `xml:"title"`
		Description    string   `xml:"description"`
		Link           string   `xml:"link"`
		PubDate        string   `xml:"pubDate"`
		LastBuildDate  string   `xml:"lastBuildDate"`
		TTL            string   `xml:"ttl"`
		Language       string   `xml:"language"`
		ManagingEditor string   `xml:"managingEditor"`
		WebMaster      string   `xml:"webMaster"`
		Image          image    `xml:"image"`
		Item           []item   `xml:"item"`
	}

	// rssDocument defines the fields associated with the rss document.
	rssDocument struct {
		XMLName xml.Name `xml:"rss"`
		Channel channel  `xml:"channel"`
	}
)

type rssMatcher struct{}

func init() {
	var matcher rssMatcher
	search.Register("rss", matcher)
}

func (m rssMatcher) Search(feed *search.Feed, searchTerm string) ([]*search.Result, error) {
	var results []*search.Result

	rssDoc, err := m.retrieve(feed)
	if err != nil {
		return nil, err
	}

	for _, channelItem := range rssDoc.Channel.Item {
		matched, err := regexp.MatchString(searchTerm, channelItem.Title)

		if err != nil {
			return nil, err
		}

		if matched {
			results = append(results, &search.Result{
				Field: "Title",
				Content: channelItem.Title,
			})
		}

		matched, err = regexp.MatchString(searchTerm, channelItem.Description)
		if err != nil {
			return nil, err
		}

		if matched {
			results = append(results, &search.Result{
				Field: "Description",
				Content: channelItem.Description,
			})
		}
	}
	return results, nil
}

func (m rssMatcher) retrieve(feed *search.Feed) (*rssDocument, error) {
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
