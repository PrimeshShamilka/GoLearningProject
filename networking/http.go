package networking

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"encoding/json"
)

type Job struct {
	User   string `json:"user"`
	Action string `json:"action"`
	Count  int    `json:"count"`
}

func main2() {
	// GET request
	resp, err := http.Get("https://example.com")
	if err != nil {
		log.Fatal("Error fetching URL:", err)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)

	// POST request
	job := &Job{
		User:   "john_doe",
		Action: "create",
		Count:  1,
	}
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(job); err != nil {
		log.Fatal("Error encoding JSON:", err)
	}

	resp2, err := http.Post("https://example.com/api", "application/json", &buf)
	if err != nil {
		log.Fatal("Error making POST request:", err)
	}
	defer resp2.Body.Close()
	io.Copy(os.Stdout, resp2.Body)
}
