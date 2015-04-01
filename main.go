package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

const activityPath = "/Users/jaffee/go/src/github.com/jaffee/github/"

type RepoActivity struct {
	Name    string
	Commits []CommitDiff
}

type CommitDiff struct {
	Metadata Commit
	Diff     string
}

type Commit struct {
	Url     string
	Message string
}

type Post struct {
	Title    string
	Sections []RepoActivity
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func serveRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><body>Letsssssss goooooooo! %s\n</body></html>", r.URL.Path[1:])
	// TODO read *.activity to see available days
}

func serveDay(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Path[len("/day/"):]
	fname := activityPath + date + ".activity"
	fmt.Println(fname)
	fbytes, err := ioutil.ReadFile(fname)
	var repoActivities []RepoActivity
	err = json.Unmarshal(fbytes, &repoActivities)

	check(err)

	// curtime := time.Now()
	// year, month, day := curtime.Date()
	// loc, _ := time.LoadLocation("Local")
	// begOfDay := time.Date(year, month, day-4, 0, 0, 0, 0, loc)
	// RepoSlice := make([]Repo, len(repos))
	// for i := range repos {
	// 	RepoSlice[i].Name = repos[i]
	// 	RepoSlice[i].Commits = github.GetCommits(username, repos[i], begOfDay)
	// }

	post := &Post{Title: date, Sections: repoActivities}
	t, _ := template.ParseFiles("realday.html")
	t.Execute(w, post)
}

func main() {
	// TODO spin off file checker goroutine to see if new data needs to be fetched
	http.HandleFunc("/", serveRoot)
	http.HandleFunc("/day/", serveDay)
	http.ListenAndServe(":8080", nil)
}
