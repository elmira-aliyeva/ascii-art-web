package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"

	ascii ".."
)

var err500 = "500 Internal Server Error :("
var err404 = "404 Oops, this page not found..."
var err400 = "400 Bad request, I can't print this!"

// Page holds font chosen by user, their input, output from asciify, error
type Page struct {
	Font, Body, Output, Error string
}

var page Page

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, 404)
		return
	}
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, page)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("body")
	body = strings.ReplaceAll(body, "\r\n", "\\n")
	if !isValid(body) {
		errorHandler(w, r, 400)
		return
	}
	page.Font = r.FormValue("fonts")
	page.Body = body
	out, err := ascii.Asciify(body, page.Font)
	if err != nil {
		fmt.Println("Error in asciify:", err)
		errorHandler(w, r, 500)
		return
	}
	page.Output = out
	http.Redirect(w, r, "/", http.StatusFound)
}

func isValid(s string) bool {
	for _, letter := range s {
		if letter < 32 || letter > 126 {
			return false
		}
	}
	return true
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	t, _ := template.ParseFiles("error.html")
	p := &Page{}
	if status == 404 {
		p = &Page{Error: err404}
	} else if status == 500 {
		p = &Page{Error: err500}
	} else if status == 400 {
		p = &Page{Error: err400}
	}
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/save/", saveHandler)

	fs := http.FileServer(http.Dir("CSS"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
