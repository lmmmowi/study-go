package goroutine

import (
	"fmt"
	"time"
)

func Channel() {
	ch := make(chan string)
	go sendData(ch)
	go getData(ch)
	time.Sleep(1e9)
}

func sendData(ch chan string) {
	ch <- "data 1"
	ch <- "data 2"
	ch <- "data 3"
	ch <- "data 4"
}

func getData(ch chan string) {
	var input string
	for {
		input = <-ch
		fmt.Println(input)
	}
}
