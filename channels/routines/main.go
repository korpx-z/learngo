package main

import (
	"fmt"
	"net/http"
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

	for {
		go checkLink(<-c, c)
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
