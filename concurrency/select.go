package concurrency

import (
	"fmt"
	"time"
)

// SELECT let's us work with multiple channels at the same time
func main() {
	ch1, ch2 := make(chan int), make (chan int)

	go func() {
		ch1 <- 42
	}()

	select {
	case val := <-ch1:
		fmt.Printf("Received from ch1: %d\n", val)
	case val := <-ch2:
		fmt.Printf("Received from ch2: %d\n", val)
	}

	// select is used for timeouts
	out := make(chan int)
	go func() {
		time.Sleep(100*time.Millisecond)
		out <- 100
	}()

	select {
	case val := <-out:
		fmt.Printf("Received from out: %d\n", val)
	case <-time.After(50 * time.Millisecond):
		fmt.Println("Timeout waiting for out")
	}
}