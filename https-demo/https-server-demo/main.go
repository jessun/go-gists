package main

import (
	"fmt"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	fmt.Println("Prepare Listen 80")
	http.ListenAndServeTLS(":80", "server.crt", "server.key", nil)
}
