package main

import (
	"fmt"
)

//as long as a type calls getGreeting() and returns a string.. it will be considered type bot (interface) as well.
type bot interface {
	getGreeting() string
}

type englishBot struct {
	greeting string
}
type spanishBot struct {
	greeting string
}

func main() {
	//	var eb englishBot
	//	eb.greeting = "hello"
	//	fmt.Printf("%v", eb.getGreeting())
	//	var eb englishBot
	//	adjustGreeting(eb)
	//	printGreeting(eb)

	var eb englishBot
	var sb spanishBot
	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Printf(b.getGreeting())
}

//func adjustGreeting(b *bot) {
//	bt := reflect.TypeOf(b).Kind()
//
//	if bt == reflect.englishBot {
//		b.greeting = "hello"
//	} else {
//		b.greeting = "hola"
//	}
//}
func (eb englishBot) getGreeting() string {
	return "hello"
}
func (sb spanishBot) getGreeting() string {
	return "hola"
}
