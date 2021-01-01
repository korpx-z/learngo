package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("gopath:", os.Getenv("GOPATH"))
}
