package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	path := args[0]

	fmt.Println(path)
}
