package main

import "reflect"

func main() {
	x := 3
	xt := reflect.TypeOf(x).Kind()

	if xt == reflect.Int {
		println(">> x is int")
	}
	if xt == reflect.Float32 {
		println(">> y is float32")
	}
	if xt == reflect.String {
		println(">> z is string")
	}
}
