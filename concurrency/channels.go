package concurrency

import (
	"fmt"
	"time"
)

/** NOTE
	channels are the preferred way to communicate between goroutines
	think as a one directional pipe which has a specific type
	receiver <- channel <- sender
	you get blocked if you try to send on a channel that has no receiver, vice versa
**/

func Main() {
	/**
		UNBUFFERED CHANNEL
		==================
	**/
	
	ch := make(chan int)
	// SEND
	// ch <- 42 // this will result in a deadlock and panic

	go func() { // go routine will run concurrently and recieve
		ch <- 42
	}()

	// RECEIVE
	val := <-ch
	fmt.Println(val)

	// SEND MULTIPLE
	count := 3
	go func() {
		// this goroutine is blocked until the main goroutine receives
		for i := 0; i < count; i++ {
			fmt.Printf("Sending: %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch) // signal we're done
	}()

	//main is blocked until the goroutine is sending
	// for i := 0; i < count; i++ {
	// 	val := <-ch
	// 	fmt.Printf("Received: %d\n", val)
	// }
	for i := range ch {
		fmt.Printf("Received: %d\n", i)
	}


	/**
		BUFFERED CHANNEL
		================
	**/

	// SEND 
	ch2 := make(chan int, 1) // buffer size of 1
	ch2 <- 23
	val2 := <-ch2
	fmt.Printf("Received: %d\n", val2)

	fmt.Println("---------------------------")
	// SEND MULTIPLE 
	ch3 := make(chan int, 4) // buffer size of 1
	count2 := 4
	for i := 0; i < count2; i++ {
		fmt.Printf("Sending: %d\n", i)
		ch3 <- i
	}
	close(ch3)

	for i := range ch3 {
		fmt.Printf("Received: %d\n", i)
	}
}
