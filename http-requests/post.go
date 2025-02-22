package main

import (
	"bytes"
	"log"
	"net/http"
	"time"
)

func main() {
	jsonStr := []byte(`{ "test": "event"}`)

	req, err := http.NewRequest("POST", "https://eotoaq34amuta9f.m.pipedream.net", bytes.NewReader(jsonStr))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: time.Minute}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	log.Printf("Status: %s", resp.Status)
	log.Printf("Body: %s", resp.Body)

	defer resp.Body.Close()
}
