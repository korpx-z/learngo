package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWrite struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Printf("%+v", err)
		os.Exit(1)
	}

	//	bs := make([]byte, 99999)
	//	resp.Body.Read(bs)
	//	fmt.Printf("%v", string(bs))
	lw := logWrite{}
	io.Copy(lw, resp.Body)
}

func (logWrite) Write(bs []byte) (int, error) {
	fmt.Printf("%v", string(bs))
	fmt.Printf("Just wrote this many bites: %v", len(bs))
	return len(bs), nil
}
