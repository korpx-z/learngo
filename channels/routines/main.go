package main

import (
	"fmt"
	"net/http"
)

func main() {

	var links []string
	links = append(links,
		"http://www.google.com",
		"https://www.ibm.com",
	)

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}
	fmt.Printf("%v"+" "+"is up! \n", link)
}
