package main

import (
	"os"
	"fmt"
	"flag"
)

var para_a = flag.Bool("a", false, "boolean parameter")
var para_b = flag.String("b", "para_b_str", "string parameter")

func main() {
	flag.Parse()
	fmt.Println("para_a=", *para_a)
	fmt.Println("para_b=", *para_b)
	fmt.Println("other_paras=", flag.Args())
	for index, arg := range os.Args {
		fmt.Println("args[", index, "] =", arg)
	}
}
