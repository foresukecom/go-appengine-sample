package main

import (
	"net/http"
	"fmt"
)

func init() {
	http.HandleFunc("/hoge", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Google App Engine for Go world.")
}



