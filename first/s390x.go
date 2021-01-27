package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("shell:", os.Getenv("SHELL"))
    fmt.Println("Built on x86, for s390x architecture using go build")
}
