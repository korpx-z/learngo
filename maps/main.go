package main

import "fmt"

func main() {
	//	colors := map[string]string{
	//		"red":   "#ff0000",
	//		"green": "4bf745",
	//	}

	colors := make(map[string]string)
	colors["red"] = "#ff0000"
	colors["green"] = "4bf745"
	colors["test"] = "DELETEME"
	delete(colors, "test")
	fmt.Printf("%+v", colors)
	iterateMap(colors)
}

func iterateMap(c map[string]string) {
	for k, v := range c {
		fmt.Printf("%[1]v"+" "+"%[2]v", k, v)
	}
}
