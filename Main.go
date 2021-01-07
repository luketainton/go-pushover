package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

// MessageResult is the response from the API
type MessageResult struct {
	Status    int    `json:"status"`
	RequestID string `json:"request"`
}

func main() {
	var title string
	var message string
	var url string
	var urltitle string
	var html bool

	flag.StringVar(&title, "t", "", "Message title")
	flag.StringVar(&message, "m", "", "Message body")
	flag.StringVar(&url, "u", "", "Hyperlink URL (if required)")
	flag.StringVar(&urltitle, "r", "", "Hyperlink text (if required)")
	flag.BoolVar(&html, "w", false, "Enable HTML parsing")
	flag.Parse()

	if title == "" || message == "" {
		fmt.Println("FATAL: Please specify message title and body.")
		os.Exit(1)
	} else {
		resultStr := sendMessage(title, message, html, url, urltitle)
		var result MessageResult
		err := json.Unmarshal([]byte(resultStr), &result)
		if err != nil {
			fmt.Println(resultStr)
		}
		fmt.Println("Request ID: " + result.RequestID)
		if result.Status == 1 {
			fmt.Println("Successful: YES")
		} else if result.Status == 0 {
			fmt.Println("Successful: NO")
		}
	}
}
