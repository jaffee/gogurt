package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

//const activityPath = "/Users/jaffee/go/src/github.com/jaffee/github/"
//const templateloc = "/Users/jaffee/go/src/github.com/jaffee/gogurt/"
//const staticloc = "/Users/jaffee/go/src/github.com/jaffee/gogurt/"
var config *Config

type Config struct {
	ActivityPath string
	Templateloc  string
	Staticloc    string
}

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

type RootPage struct {
	Dates []string
}

func serveRoot(w http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir(config.ActivityPath)
	check(err)
	dates := make([]string, len(files))
	for i, f := range files {
		name := f.Name()
		if strings.HasSuffix(name, ".activity") {
			loc := strings.LastIndex(name, ".activity")
			dates[i] = name[:loc]
		}
	}
	rootpg := &RootPage{Dates: dates}
	t, err := template.ParseFiles(config.Templateloc + "root.html")
	check(err)
	t.Execute(w, rootpg)
}

func serveDay(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Path[len("/day/"):]
	fname := config.ActivityPath + date + ".activity"
	fmt.Println(fname)
	fbytes, err := ioutil.ReadFile(fname)
	var repoActivities []RepoActivity
	err = json.Unmarshal(fbytes, &repoActivities)

	check(err)
	post := &Post{Title: date, Sections: repoActivities}
	t, err := template.ParseFiles(config.Templateloc + "day.html")
	check(err)
	t.Execute(w, post)
}

func serveStatic(w http.ResponseWriter, r *http.Request) {
	fname := config.Staticloc + r.URL.Path[len("/static/"):]
	http.ServeFile(w, r, fname)
}

func main() {
	conff, err := ioutil.ReadFile("config.json")
	fmt.Printf("AAA%v\n", conff)

	check(err)
	config = &Config{}
	err = json.Unmarshal(conff, config)
	check(err)

	http.HandleFunc("/", serveRoot)
	http.HandleFunc("/day/", serveDay)
	http.HandleFunc("/static/", serveStatic)
	http.ListenAndServe(":8080", nil)
}
