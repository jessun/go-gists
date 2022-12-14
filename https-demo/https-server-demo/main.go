package main

import (
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/soheilhy/cmux"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}

func Https1() {
	http.HandleFunc("/hello", HelloServer)
	fmt.Println("Prepare Listen 8080")
	http.ListenAndServeTLS(":8080", "server.crt", "server.key", nil)
}

func main() {

	b, err := ioutil.ReadFile("./server.crt")
	if err != nil {
		panic(err)
	}

	p, _ := pem.Decode(b)
	crt, err := x509.ParseCertificate(p.Bytes)
	if err != nil {

	}
	fmt.Printf("crt<%#v>\n", crt.NotAfter.Local().Format(time.RFC3339))
	fmt.Printf("crt<%#v>\n", crt.NotBefore.Local().Format(time.RFC3339))
	Https1()
	// Example_bothHTTPAndHTTPS()
}

type anotherHTTPHandler struct{}

func (h *anotherHTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "example http response")
}

func serveHTTP1(l net.Listener) {
	s := &http.Server{
		Handler: &anotherHTTPHandler{},
	}
	if err := s.Serve(l); err != cmux.ErrListenerClosed {
		panic(err)
	}
}

func serveHTTPS(l net.Listener) {
	// Load certificates.
	certificate, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Panic(err)
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{certificate},
		Rand:         rand.Reader,
	}

	// Create TLS listener.
	tlsl := tls.NewListener(l, config)

	// Serve HTTP over TLS.
	serveHTTP1(tlsl)
}

// This is an example for serving HTTP and HTTPS on the same port.
func Example_bothHTTPAndHTTPS() {
	// Create the TCP listener.
	l, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Panic(err)
	}

	// Create a mux.
	m := cmux.New(l)

	// We first match on HTTP 1.1 methods.
	// httpl := m.Match(cmux.HTTP1Fast())

	// If not matched, we assume that its TLS.
	//
	// Note that you can take this listener, do TLS handshake and
	// create another mux to multiplex the connections over TLS.
	tlsl := m.Match(cmux.Any())

	// go serveHTTP1(httpl)
	go serveHTTPS(tlsl)

	if err := m.Serve(); !strings.Contains(err.Error(), "use of closed network connection") {
		panic(err)
	}
}
