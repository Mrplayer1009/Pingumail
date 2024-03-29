package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Mail struct {
	ID   int    `json:"id"`
	Body string `json:"body"`
}

// Set content of mails to the content of mails.json
var mails []Mail

func init() {
	content, err := ioutil.ReadFile("pingumail.json")
	if err != nil {
		fmt.Println("Error reading mails from file:", err)
		return
	}

	err = json.Unmarshal(content, &mails)
	if err != nil {
		fmt.Println("Error unmarshaling mails from JSON:", err)
		return
	}
}

func Start() {

	println("Mails loaded from file:", mails)

	http.HandleFunc("/mail", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			json.NewEncoder(w).Encode(mails)
		case http.MethodPost:
			var mail Mail
			if err := json.NewDecoder(r.Body).Decode(&mail); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			mail.ID = len(mails)

			mails = append(mails, mail)
			json.NewEncoder(w).Encode(mail)

			// Write mails to JSON file
			jsonData, err := json.Marshal(mails)
			if err != nil {
				fmt.Println("Error marshaling mails to JSON:", err)
				return
			}

			err = ioutil.WriteFile("pingumail.json", jsonData, 0644)
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
