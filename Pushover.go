package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

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
