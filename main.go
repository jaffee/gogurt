package main

import (
	"fmt"
	"github.com/jaffee/gogurt/github"
	"html/template"
	"net/http"
	"strings"
	"time"
)

type Post struct {
	Title    string
	Sections []Repo
}

type Repo struct {
	Name    string
	Commits []github.Commit
}

func (r *Repo) String() string {
	var cstrings []string
	for i := range r.Commits {
		cstrings := append(cstrings, r.Commits[i].String())
	}
	return r.Name + "/n" + strings.Join(cstrings, "\n")
}

func serveRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><body>Letsssssss goooooooo! %s\n</body></html>", r.URL.Path[1:])
}

func serveDay(w http.ResponseWriter, r *http.Request) {
	repos := [...]string{"gogurt", "robpike.io", "goplait"}
	username := "jaffee"
	date := r.URL.Path[len("/day/"):]
	curtime := time.Now()
	year, month, day := curtime.Date()
	loc, _ := time.LoadLocation("Local")
	begOfDay := time.Date(year, month, day-4, 0, 0, 0, 0, loc)
	RepoSlice := make([]Repo, len(repos))
	for i := range repos {
		RepoSlice[i].Name = repos[i]
		RepoSlice[i].Commits = github.GetCommits(username, repos[i], begOfDay)
	}

	post := &Post{Title: date, Sections: RepoSlice}
	t, _ := template.ParseFiles("day.html")
	t.Execute(w, post)
}

func main() {
	http.HandleFunc("/", serveRoot)
	http.HandleFunc("/day/", serveDay)
	http.ListenAndServe(":8080", nil)
}
