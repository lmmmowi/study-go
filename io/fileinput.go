package io

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func FileInput() {
	filePath := "assets/test.md"
	inputFile, openError := os.Open(filePath)
	if openError != nil {
		fmt.Println("Fail to open file")
		return
	}
	defer inputFile.Close()

	// by ReadString
	reader := bufio.NewReader(inputFile)
	for {
		s, readError := reader.ReadString('\n')
		fmt.Print(s)
		if readError == io.EOF {
			break
		}
	}

	// by Read
	fileLength, _ := inputFile.Seek(0, io.SeekEnd)
	inputFile.Seek(0, io.SeekStart)
	buf := make([]byte, fileLength)
	reader.Read(buf)
	s := string(buf[:])
	fmt.Println(s)

	// by Fscan
	inputFile.Seek(0, io.SeekStart)
	var col1, col2, col3 []string
	for {
		var s1, s2, s3 string
		_, error := fmt.Fscanln(reader, &s1, &s2, &s3)
		if error != nil {
			break
		}
		col1 = append(col1, s1)
		col2 = append(col2, s2)
		col3 = append(col3, s3)
	}
	fmt.Println(col1)
	fmt.Println(col2)
	fmt.Println(col3)
}
