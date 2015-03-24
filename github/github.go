package github

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
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

func GetCommits(username string, since time.Time) []string {
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
	var all_commits []byte
	for i := 0; i < len(repos_slice); i++ {
	 	all_commits = append(all_commits, getBody(repos_slice[i].Commits_url + "?since=" + since.Format(time_layout))...)
	}
	all_strs := make([]string, len(all_commits))
	for i := 0; i < len(all_commits); i++ {
	 	all_strs[i] = string(all_commits[i])
	}
	all_strs := []string{"a", "b"}
	return all_strs
}

func build_commit_url(com_url string, since string) {
	// get rid of {sha} from com_url and add this since parameter
	// use this method in the all_commits = append... line above
}

func getBody(url string) []byte {
	resp, err := http.Get(url)
	checkErr(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)
	return body
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
