package io

import (
	"io"
	"os"
)

func FileCopy() {
	src := "assets/test.md"
	dest := "assets/test-copy.md"

	inputFile, _ := os.Open(src)
	defer inputFile.Close()

	_, err := os.Lstat(dest)
	if !os.IsNotExist(err) {
		os.Remove(dest)
	}

	outputFile, _ := os.Create(dest)
	defer outputFile.Close()

	io.Copy(outputFile, inputFile)
}
