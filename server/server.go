package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"os"

	"golang.org/x/crypto/bcrypt"
	_ "golang.org/x/term"
)

const jsonPath = "pingumail.json"

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
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Token    string `json:"token"`
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
				if mail.To == os.Getenv("pinguUserName") && !mail.Read {
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

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			json.NewEncoder(w).Encode(jsonBDD.USers)
		default:
			http.Error(w, "Status Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:

			// Read the username from the request and the password
			var user User
			if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
				http.Error(w, "Status Bad Request", http.StatusBadRequest)
				return
			}

			test, _ := json.Marshal(user)
			json.NewEncoder(w).Encode(string(test))

			for _, u := range jsonBDD.USers {
				if u.Name == user.Name {
					err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password))

					json.NewEncoder(w).Encode(u.Name)

					if err == nil {
						fmt.Println("Login successful")
						json.NewEncoder(w).Encode(u.Token)
						return
					}
				}
			}

		default:
			http.Error(w, "Status Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server started on port 80 (HTTP)")
	if err := http.ListenAndServe(":80", nil); err != nil {
		fmt.Println(err)
	}
}

func AddUser(name string, password string) {
	var user User
	user.ID = len(jsonBDD.USers) + 1
	user.Name = name
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error hashing password")
		return
	}
	user.Password = string(hash)

	// if the user already exists
	for _, u := range jsonBDD.USers {
		if u.Name == name {
			fmt.Printf("User %s already exists\n", name)
			return
		}
	}

	jsonBDD.USers = append(jsonBDD.USers, user)

	jsonData, err := json.Marshal(jsonBDD)
	handleErr(err, "Error marshaling mails to JSON")

	err = ioutil.WriteFile(jsonPath, jsonData, 0644)
	handleErr(err, "Error writing mails to file")

	fmt.Printf("User %s added\n", name)
}
