package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Mail struct {
	ID   int    `json:"id"`
	From string `json:"from"`
	To   string `json:"to"`
	Body string `json:"body"`
}

const server = "http://localhost:80/mail"

func handleErr(err error, reason string) {
	if err != nil {
		fmt.Println("Error:", reason, err)
	}
}

func SendMail(from string, to string, body string) {
	var mail = Mail{
		From: from,
		To:   to,
		Body: body,
	}

	bodyRequest, err := json.Marshal(mail) // Replace with your custom body
	handleErr(err, "Error marshalling body")

	req, err := http.NewRequest("POST", server, bytes.NewBuffer(bodyRequest))
	handleErr(err, "Error creating request")

	client := &http.Client{}
	resp, err := client.Do(req)
	handleErr(err, "Error making request")
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	handleErr(err, "Error reading response body")

	fmt.Println("Response:", string(responseBody))
}
