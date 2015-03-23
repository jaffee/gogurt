package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type user struct {
	Login     string
	Repos_url string
}

// type repos struct {
	
// }

// type repo struct {
	
// }



func main() {
	resp, _ := http.Get("https://api.github.com/users/jaffee")
	fmt.Printf("%v\n", resp)
	fmt.Printf("%v\n", resp.Body)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var u user	
	err := json.Unmarshal(body, &u)
	checkErr(err)
	fmt.Printf("%v\n", u)
	
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
