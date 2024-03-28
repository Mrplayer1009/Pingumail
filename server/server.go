package server

import (
	"fmt"
	"net/http"
)

func Start() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "[%s] : %s\n", r.Method, r.URL.Path)
	})

	fmt.Println("Server started on port 80 (HTTP)")
	if err := http.ListenAndServe(":80", nil); err != nil {
		fmt.Println(err)
	}

}