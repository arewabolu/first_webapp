package main

import (
	"net/http"
)

func main() {
	multi := http.NewServeMux() // Create a new Server multiplexer using ServeMux
	multi.HandleFunc("/index", index)
	multi.HandleFunc("/signup", signup)
}

func index(w http.ResponseWriter, T *http.Request) {
	T.URL.Parse()
}

func signup() {

}
