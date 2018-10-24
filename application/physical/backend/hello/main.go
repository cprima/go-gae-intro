package main

import (
	"fmt"
	"net/http"
	"strings"

	"google.golang.org/appengine"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {

	name := strings.Split(strings.TrimPrefix(
		r.URL.Path, "/hello"),
		"/")[1]

	if name != "" {
		fmt.Fprintf(w, "Hello %v!", name)
	} else {
		fmt.Fprintln(w, "Hello world!")
	}
}

func main() {
	http.HandleFunc("/hello/", helloHandler)
	appengine.Main()
}
