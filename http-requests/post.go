package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
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
	// Read the response body
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Unmarshal the JSON
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, bodyBytes, "", "  "); err != nil {
		panic(err)
	}

	// Print the formatted JSON
	log.Printf("Body: %s", prettyJSON.String())

	defer resp.Body.Close()
}
