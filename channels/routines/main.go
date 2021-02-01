package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	c := make(chan string)

	var links []string
	links = append(links,
		"http://www.google.com",
		"https://www.ibm.com",
	)

	for _, link := range links {
		go checkLink(link, c)
	}

	//    for {
	//		go checkLink(<-c, c)
	//	}
	// refactored vvvv
	for l := range c {
		go func(link string) {
			time.Sleep(2 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("%+v", err)
		c <- link
	}
	fmt.Printf("%v"+" "+"is up! \n", link)
	c <- link
}
