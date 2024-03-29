package server

import (
	"encoding/json"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"io/ioutil"
	"net/http"
)

var jsonPath = "pingumail.json"

type BDD struct {
	Mails []Mail `json:"mails"`
	Users []User `json:"users"`
}

type Mail struct {
	ID   int    `json:"id"`
	Body string `json:"body"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Set content of mails to the content of mails.json
var jsonBDD BDD

func init() {
	content, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		fmt.Println("Error reading mails from file:", err)
		return
	}

	err = json.Unmarshal(content, &jsonBDD.Mails)
	if err != nil {
		fmt.Println("Error unmarshaling mails from JSON:", err)
		return
	}
}

func Start() {

	println("Mails loaded from file:", jsonBDD.Mails)

	http.HandleFunc("/mail", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			json.NewEncoder(w).Encode(jsonBDD.Mails)
		case http.MethodPost:
			var mail Mail
			if err := json.NewDecoder(r.Body).Decode(&mail); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			mail.ID = len(jsonBDD.Mails)

			jsonBDD.Mails = append(jsonBDD.Mails, mail)
			json.NewEncoder(w).Encode(mail)

			// Write jsonBDD.Mails to JSON file
			jsonData, err := json.Marshal(jsonBDD)
			if err != nil {
				fmt.Println("Error marshaling mails to JSON:", err)
				return
			}

			err = ioutil.WriteFile(jsonPath, jsonData, 0644)
			if err != nil {
				fmt.Println("Error writing mails to file:", err)
				return
			}

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server started on port 80 (HTTP)")
	if err := http.ListenAndServe(":80", nil); err != nil {
		fmt.Println(err)
	}

}

