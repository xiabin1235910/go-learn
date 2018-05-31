package main

import (
	"fmt"
	"net/http"
	"time"
)

func main () {
	
	links := []string {
		"http://www.baidu.com",
		"http://www.taobao.com",
		"http://www.jd.com",
	};

	fmt.Println(links)
	fmt.Println("hello world")

	c := make(chan string)
	
	// for _, link := range links {
	// 	// fmt.Println(i)
		
	// 	go checkLink(link, c)
	// //	fmt.Println(<- c)
	// }

	//for {
	//	fmt.Println(<- c)
	//}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
		//go checkLink(l, c)
	}
	
	// checkLink();
}

func checkLink(link string, c chan string) {
	//time.Sleep(5 * time.Second);
	
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("url is wrong")
	} else {
		c <- link
		fmt.Println(link);
	}
}
