package concurrency

import (
	"context"
	"fmt"
	"time"
)

/** NOTE
	context is a package used for deadlines and cancellations
**/

type Bid struct {
	AdUrl string
	Price float64
}

func bestBid(url string) Bid {
	time.Sleep(200 * time.Millisecond) // simulate network call

	return Bid{
		AdUrl: "http://ads.com/",
		Price: 1.23,
	}
}

var defaultBid = Bid{
	AdUrl: "http://ads1.com/",
	Price: 0.50,
}

func findBid(ctx context.Context, url string) Bid {
	ch := make(chan Bid, 1) // buffered channel to avoid goroutine leak
	go func(){
		ch <- bestBid(url)
	}()

	select {
	case bid := <-ch:
		return bid
	case <-ctx.Done():
		fmt.Println("context cancelled or timed out:", ctx.Err())
		return defaultBid
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel() // good practice to call cancel in case the function returns early
	bid := findBid(ctx, "http://example.com/")
	fmt.Printf("Found bid: %+v\n", bid) // prints best bid

	ctx2, cancel2 := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel2()
	bid2 := findBid(ctx2, "http://example.com/")
	fmt.Printf("Found bid: %+v\n", bid2) // prints default bid due to timeout
}

