package main

import (
	"os"
	"fmt"
)

func main() {
	for index, arg := range os.Args {
		fmt.Println("args[", index, "] =", arg)
	}
}
