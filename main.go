package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Post struct {
	Title string
}

func serveRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><body>Letsssssss goooooooo! %s\n</body></html>", r.URL.Path[1:])
}

func serveDay(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Path[len("/day/"):]
	post := &Post{Title: date}
	t, _ := template.ParseFiles("day.html")
	t.Execute(w, post)
}

func main() {
	http.HandleFunc("/", serveRoot)
	http.HandleFunc("/day/", serveDay)
	http.ListenAndServe(":8080", nil)
}
