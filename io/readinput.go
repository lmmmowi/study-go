package io

import (
	"bufio"
	"fmt"
	"os"
)

func ReadInput() {
	var name string

	fmt.Println("Pleas input your name")
	fmt.Scanln(&name)
	fmt.Printf("Hello, %s", name)

	fmt.Println("Pleas input your name again")
	reader := bufio.NewReader(os.Stdin)
	name, _ = reader.ReadString('\n')
	fmt.Printf("Hello, %s", name)
}
