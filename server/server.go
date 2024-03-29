package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const jsonPath = "pingumail.json"
const currentUser = "mathis"

type BDD struct {
	Mails []Mail `json:"mails"`
	USers []User `json:"users"`
}

type Mail struct {
	ID   int    `json:"id"`
	From string `json:"from"`
	To   string `json:"to"`
	Body string `json:"body"`
	Read bool   `json:"read"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Set content of mails to the content of mails.json
var jsonBDD BDD

func handleErr(err error, reason string) {
	if err != nil {
		fmt.Println("Error:", reason)
	}
}

func init() {
	content, err := ioutil.ReadFile(jsonPath)
	handleErr(err, "reading mails from file")

	err = json.Unmarshal(content, &jsonBDD)
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

			var mails []Mail
			var backupMail []Mail

			for _, mail := range jsonBDD.Mails {
				if mail.To == currentUser && !mail.Read {
					mail.Read = true
					mails = append(mails, mail)
				}
				backupMail = append(backupMail, mail)
			}

			jsonBDD.Mails = backupMail

			jsonData, err := json.Marshal(jsonBDD)
			handleErr(err, "Error marshaling mails to JSON")

			err = ioutil.WriteFile(jsonPath, jsonData, 0644)
			handleErr(err, "Error writing mails to file")

			json.NewEncoder(w).Encode(mails)

		case http.MethodPost:
			var mail Mail
			if err := json.NewDecoder(r.Body).Decode(&mail); err != nil {
				http.Error(w, "Status Bad Request", http.StatusBadRequest)
				return
			}

			mail.ID = len(jsonBDD.Mails)
			mail.Read = false

			jsonBDD.Mails = append(jsonBDD.Mails, mail)
			json.NewEncoder(w).Encode(mail)

			// Write mails to JSON file
			jsonData, err := json.Marshal(jsonBDD)
			handleErr(err, "Error marshaling mails to JSON")

			err = ioutil.WriteFile(jsonPath, jsonData, 0644)
			handleErr(err, "Error writing mails to file")

			fmt.Printf("Sending mail from to %s\n", mail.To)

		default:
			http.Error(w, "Status Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server started on port 80 (HTTP)")
	if err := http.ListenAndServe(":80", nil); err != nil {
		fmt.Println(err)
	}

}
