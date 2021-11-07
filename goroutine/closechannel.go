package goroutine

import (
	"fmt"
	"time"
)

func CloseChannel() {
	ch := make(chan int)
	go sendDataAndClose(ch)

	for {
		v, ok := <-ch
		if ok {
			fmt.Println(v)
		} else {
			break
		}
	}
}

func sendDataAndClose(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	time.Sleep(1e9)
	ch <- 4
	ch <- 5
	close(ch)
}
