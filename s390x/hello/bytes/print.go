package main

import "fmt"

var printBytes string

func main() {

	//x86 ASCII prints [104 101 108 108 111]
	//s390x EBCDIC prints
	printBytes := "hello"
	fmt.Println([]byte(printBytes))
}
