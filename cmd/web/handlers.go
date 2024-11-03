package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("server", "go")
	// w.Write([]byte("Hello from snippetbox"))

	templ, err := template.ParseFiles("./ui/html/base.tmpl.html", "./ui/html/pages/home.tmpl.html")

	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error 1", http.StatusInternalServerError)
		return
	}

	err = templ.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// msg := fmt.Sprintf("Display a specific snippet with id %d...", id)
	// w.Write([]byte(msg))
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet...."))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Creating and saving a new snippet...."))
}
