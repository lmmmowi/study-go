package goroutine

import (
	"fmt"
	"time"
)

func Select() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go sendDataAlways(ch1)
	go sendDataAlways(ch2)
	go receiveData(ch1, ch2)
	time.Sleep(1e9)
}

func sendDataAlways(ch chan int) {
	for i := 0; true; i++ {
		ch <- i
		time.Sleep(1e6)
	}
}

func receiveData(ch1, ch2 chan int) {
	for {
		select {
		case v1 := <-ch1:
			fmt.Printf("receive data on channel1: %d\n", v1)
		case v2 := <-ch2:
			fmt.Printf("receive data on channel2: %d\n", v2)
		}
	}
}
