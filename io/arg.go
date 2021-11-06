package io

import (
	"flag"
	"fmt"
)

// go-study.exe -n A B C
func HandleArg() {
	newLine := flag.Bool("n", false, "print newline")

	flag.PrintDefaults()
	flag.Parse()

	s := ""
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 && *newLine {
			s += "\n"
		}
		s += flag.Arg(i)
	}
	fmt.Println(s)
}
