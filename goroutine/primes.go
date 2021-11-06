package goroutine

import "fmt"

func GetPrimes() {
	primes := generate()

	for i := 0; i < 20; i++ {
		p := <-primes
		fmt.Printf("%d ", p)

		ch := make(chan int)
		go filter(p, primes, ch)
		primes = ch
	}
}

func generate() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func filter(prime int, in chan int, out chan int) {
	for i := range in {
		if i%prime != 0 {
			out <- i
		}
	}
}
