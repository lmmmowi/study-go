package goroutine

import (
	"fmt"
	"time"
)

func Sleep() {
	go longWait()
	go shortWait()
	fmt.Println("About to sleep in main")
	time.Sleep(10 * 1e9)
	fmt.Println("End of sleep in main")
}

func longWait() {
	fmt.Println("About to sleep in longWait")
	time.Sleep(5 * 1e9)
	fmt.Println("End of sleep in longWait")
}

func shortWait() {
	fmt.Println("About to sleep in shortWait")
	time.Sleep(2 * 1e9)
	fmt.Println("End of sleep in shortWait")
}
