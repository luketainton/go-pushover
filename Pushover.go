package pushover

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// MessageResult is the response from the API
type MessageResult struct {
	Status    int    `json:"status"`
	RequestID string `json:"request"`
}

func sendMessage(title string, message string, htmlBool bool, url string, urlTitle string) string {
	var apiEndpoint string = "https://api.pushover.net/1/messages.json"
	var apiUserToken string = os.Getenv("PUSHOVER_USER_TOKEN")
	var apiAppToken string = os.Getenv("PUSHOVER_APP_TOKEN")

	if apiAppToken == "" || apiUserToken == "" {
		fmt.Println("FATAL: API app/user token not set. Please set PUSHOVER_USER_TOKEN and PUSHOVER_APP_TOKEN environment variables.")
		os.Exit(2)
	}

	var html string
	if htmlBool == true {
		html = "1"
	} else {
		html = "0"
	}

	reqBody, err := json.Marshal(map[string]string{
		"token":     apiAppToken,
		"user":      apiUserToken,
		"html":      html,
		"title":     title,
		"message":   message,
		"url":       url,
		"url_title": urlTitle,
	})

	if err != nil {
		fmt.Print("FATAL: ")
		fmt.Println(err)
		os.Exit(3)
	}

	res, err := http.Post(apiEndpoint, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Print("FATAL: ")
		fmt.Println(err)
		os.Exit(4)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Print("FATAL: ")
		fmt.Println(err)
		os.Exit(5)
	}

	return string(body)
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
