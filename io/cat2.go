package io

import (
	"flag"
	"os"
)

func RunCat2() {
	flag.Parse()

	if flag.NArg() == 0 {
		cat2(os.Stdin)
	} else {
		for i := 0; i < flag.NArg(); i++ {
			f, err := os.Open(flag.Arg(i))
			if err == nil {
				cat2(f)
			}
			f.Close()
		}
	}
}

func cat2(file *os.File) {
	buf := make([]byte, 512)
	for {
		rn, _ := file.Read(buf)
		if rn > 0 {
			os.Stdout.Write(buf[0:rn])
		} else if rn == 0 {
			break
		}
	}
}
