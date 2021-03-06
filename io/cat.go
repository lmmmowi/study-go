package io

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func RunCat() {
	flag.Parse()

	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	} else {
		for i := 0; i < flag.NArg(); i++ {
			f, err := os.Open(flag.Arg(i))
			if err == nil {
				cat(bufio.NewReader(f))
			}
			f.Close()
		}
	}
}

func cat(reader *bufio.Reader) {
	for {
		s, err := reader.ReadString('\n')
		fmt.Println(s)
		if err == io.EOF {
			break
		}
	}
}
