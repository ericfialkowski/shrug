package main

import (
	"flag"
	"fmt"
)

const shrug = "¯\\_(ツ)_/¯"

var newline = flag.Bool("n", false, "Output with New Line")

func init() {
	flag.Parse()
}

func main() {
	if *newline {
		fmt.Println(shrug)
	} else {
		fmt.Print(shrug)
	}
}
