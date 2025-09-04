package networking

/** Fallacies of Distributed Computing
	1. The Network is Reliable
	2. Latency is Zero
	3. Bandwidth is Infinite
	4. The Network is Secure

	NOTE: Hence we need to set timeouts and size limits when dealing with external services
**/

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func Main() {
	fmt.Println("Networking: Timeout and Size Limit")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Millisecond)
	defer cancel() // good practice to call cancel in case the function returns early

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://example.com", nil)
	if err != nil {
		log.Fatal("Error creating request:", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Error making request:", err)
	}
	defer resp.Body.Close()

	// Limit the size of the response body to 1MB
	const mb = 1 << 20 // 1MB
	r := io.LimitReader(resp.Body, mb)
	io.Copy(os.Stdout, r)
}
