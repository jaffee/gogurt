package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type user struct {
	Login     string
	Repos_url string
}

// type repos struct {

// }

type repo struct {
	Git_commits_url string
	Commits_url     string
}

type commit struct {
	Sha string
	Url string
}

func (com *commit) String() string {
	s, _ := json.Marshal(com)
	return string(s)
}

// func main() {
// 	body := getBody("https://api.github.com/users/jaffee")
// 	var u user
// 	err := json.Unmarshal(body, &u)
// 	checkErr(err)

// 	body = getBody(u.Repos_url)
// 	repos_slice := make([]repo, 40)
// 	json.Unmarshal(body, &repos_slice)
// 	fmt.Printf("%v\n", repos_slice)
// 	fmt.Printf("%v\n", len(repos_slice))

// 	for i := 0; i < len(repos_slice); i++ {
// 		body = getBody(repos_slice[i].Commits_url)

// 	}
// }

func GetCommitsFromUsername(username string, since time.Time) []string {
	body := getBody("https://api.github.com/users/" + username)
	var u user
	err := json.Unmarshal(body, &u)
	checkErr(err)

	body = getBody(u.Repos_url)
	repos_slice := make([]repo, 40)
	json.Unmarshal(body, &repos_slice)
	fmt.Printf("%v", repos_slice)

	time_layout := "2006-01-02T15:04:05Z"
	fmt.Printf("%v", since.Format(time_layout))
	var all_commits []commit
	for i := 0; i < len(repos_slice); i++ {
		body = getBody(build_commit_url(repos_slice[i].Commits_url, since.Format(time_layout)))
		commits_slice := make([]commit, 100)
		json.Unmarshal(body, &commits_slice)
		all_commits = append(all_commits, commits_slice...)
	}
	all_strs := make([]string, len(all_commits))
	for i := 0; i < len(all_commits); i++ {
		all_strs[i] = all_commits[i].String()
	}
	//	all_strs := []string{"a", "b"}
	return all_strs
}

func main() {
	loc, _ := time.LoadLocation("Local")
	GetCommits("jaffee", "gogurt", time.Date(2014, 1, 1, 1, 1, 1, 1, loc))
}

func GetCommits(username string, repo string, since time.Time) []string {
	base_url := "https://api.github.com/repos/" + username + "/" + repo
	time_layout := "2006-01-02T15:04:05Z"
	body := getBody(base_url + "?since=" + since.Format(time_layout))
	commits_slice := make([]commit, 1)
	json.Unmarshal(body, &commits_slice) // This doesn't seem to work
	fmt.Printf("commit slice %v\n", commits_slice)
	return stringifyCommitSlice(commits_slice)
}

func stringifyCommitSlice(all_commits []commit) []string {
	all_strs := make([]string, len(all_commits))
	for i := 0; i < len(all_commits); i++ {
		all_strs[i] = all_commits[i].String()
	}
	//	all_strs := []string{"a", "b"}
	return all_strs
}

func build_commit_url(com_url string, since string) string {
	com_url = strings.Replace(com_url, "{/sha}", "", 1)
	com_url = com_url + "?since=" + since
	return com_url
	// get rid of {sha} from com_url and add this since parameter
	// use this method in the all_commits = append... line above
}

func getBody(url string) []byte {
	resp, err := http.Get(url)
	if r := handle_url_err(url, err); r != nil {
		return r
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if r := handle_url_err(url, err); r != nil {
		return r
	}
	fmt.Printf("Got body for URL %v\n%v\n", url, string(body))
	return body
}

func handle_url_err(url string, err error) []byte {
	if err != nil {
		fmt.Printf("Error with URL: %v\n", url)
		return make([]byte, 0)
	}
	return nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
