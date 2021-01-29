package main

import (
	"fmt"
	"reflect"
)

//can not combine floats and strings. literal decl defaults their values

func main() {
	x := 4.0
	y := 2
	xy := (x / y)
	xyr := reflect.TypeOf(xy).Kind()
	fmt.Printf("%v", xy)
	fmt.Printf("%v", xyr)
	f := 3.0
	b := 2
	fb := (f / b)
	fbr := reflect.TypeOf(fb).Kind()
	fmt.Printf("%v", fb)
	fmt.Printf("%v", fbr)
}
