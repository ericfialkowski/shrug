package main

import (
	"flag"
	"fmt"
	"github.com/atotto/clipboard"
)

const shrug = "¯\\_(ツ)_/¯"
const escapedShrug = "¯\\\\_(ツ)_/¯"

var newline = flag.Bool("n", false, "Output with New Line")
var escape = flag.Bool("e", false, "Escape the output")
var clip = flag.Bool("c", true, "Copy to clipboard")
var stdOut = flag.Bool("s", true, "Print output, if false copy to clipboard is forced true")

func init() {
	flag.Parse()
}

func gen() string {
	ec := ""
	if *newline {
		ec = "\n"
	}
	s := shrug
	if *escape {
		s = escapedShrug
	}
	return fmt.Sprintf("%s%s", s, ec)
}

func main() {
	s := gen()
	if *clip || !*stdOut {
		_ = clipboard.WriteAll(s)
	}
	if *stdOut {
		fmt.Print(s)
	}
}
