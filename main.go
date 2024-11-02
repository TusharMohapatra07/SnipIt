package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	log.Println("Server started on http://localhost:4000")
	err := http.ListenAndServe(":4000", mux)

	log.Fatal(err)
}
