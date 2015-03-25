package main

import (
	"fmt"
	"github.com/jaffee/gogurt/github"
	"html/template"
	"net/http"
	"time"
)

type Post struct {
	Title   string
	Commits []string
}

func serveRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><body>Letsssssss goooooooo! %s\n</body></html>", r.URL.Path[1:])
}

func serveDay(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Path[len("/day/"):]
	curtime := time.Now()
	year, month, day := curtime.Date()
	loc, _ := time.LoadLocation("Local")
	begOfDay := time.Date(year, month, day-2, 0, 0, 0, 0, loc)
	post := &Post{Title: date, Commits: github.GetCommits("jaffee", begOfDay)}
	t, _ := template.ParseFiles("day.html")
	t.Execute(w, post)
}

func main() {
	http.HandleFunc("/", serveRoot)
	http.HandleFunc("/day/", serveDay)
	http.ListenAndServe(":8080", nil)
}
