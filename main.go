package main

import (
	"fmt"
	"net/http"
)

func serveRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Letsssssss goooooooo! %s", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", serveRoot)
	http.ListenAndServe(":8080", nil)
}
