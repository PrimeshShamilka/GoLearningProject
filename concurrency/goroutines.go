package concurrency

import (
	"fmt"
	"net/http"
	"sync"
)

// Goroutines are lightweight threads managed by the Go runtime

func returnType(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	ctype := resp.Header.Get("Content-Type")
	fmt.Printf("Content-Type for %s: %s\n", url, ctype)
}

func siteSerial(urls []string) {
	for _, url := range urls {
		returnType(url)
	}
}

func sitesConcurrent(urls []string) {
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		// use the go keyword to run this in a separate goroutine
		go func(url string) {
			returnType(url)
			wg.Done() // signal that this goroutine is done
		}(url)
	}
	wg.Wait()
}

// func Main() {
// 	urls := []string{
// 		"https://example.com",
// 		"https://golang.org",
// 		"https://pkg.go.dev",
// 	}
// 	start := time.Now()
// 	siteSerial(urls)
// 	fmt.Println(time.Since(start))

// 	start = time.Now()
// 	sitesConcurrent(urls)
// 	fmt.Println(time.Since(start))
// }
