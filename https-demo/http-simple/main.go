package main

import "net/http"

func Hello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("hello"))
}

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":10000", nil)
}
