package networking

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

// JSON encoding and decoding
type Request struct {
	Login  string  `json:"login"`
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}

type Response struct {
	Status  string  `json:"status"`
	Balance float64 `json:"balance"`
}

var data = `
{
	"login": "user1",
	"type": "withdraw",
	"amount": 100.0
}
`

func main1() {

	// Decoding JSON
	rdr := strings.NewReader(data) // simulates a file/socket
	dec := json.NewDecoder(rdr)
	var req Request
	if err := dec.Decode(&req); err != nil {
		log.Fatal("decode error:", err)
	}
	fmt.Printf("decoded: %+v\n", req)

	// Encoding JSON
	resp := Response{
		Status:  "ok",
		Balance: 900.0,
	}
	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(resp); err != nil {
		log.Fatal("encode error:", err)
	}
}
