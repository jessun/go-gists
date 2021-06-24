package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var currentServer *http.Server
var isHttpsEnableChan chan bool = make(chan bool, 1)
var isHttpsEnabled *bool

var changeHttpsEnableByUser bool

func main() {

	isHttpsEnableChan <- false
	httpServerWatcher()

}

func createServer() {
	currentServer = &http.Server{
		Addr:    ":8000",
		Handler: &myHandler{},
	}
}

type myHandler struct{}

func (c myHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read request body got error: <%v>", err)
		return
	}
	body := string(bs)

	fmt.Printf("request body: <%v>\n", body)

	if strings.Contains(body, "https enable") {
		isHttpsEnableChan <- true
		changeHttpsEnableByUser = true
	}

	if strings.Contains(body, "https disable") {
		isHttpsEnableChan <- false
		changeHttpsEnableByUser = true
	}

	fmt.Println("req handle finished.")
}

func httpServerWatcher() {

	for {
		select {
		case isHttpsNew := <-isHttpsEnableChan:
			if isHttpsEnabled == nil {
				isHttpsEnabled = &isHttpsNew
				go runServer()
			} else {
				if *isHttpsEnabled != isHttpsNew {
					fmt.Printf("change https from <%v> to <%v> \n", getIsHttpsEnabled(), isHttpsNew)
					*isHttpsEnabled = isHttpsNew
					restartServer()
				}
			}
		}
	}
}

func runServer() {
	createServer()

	go func() {
		if isHttpsEnabled != nil && *isHttpsEnabled {
			fmt.Println("run server with https enable...")
			if err := currentServer.ListenAndServeTLS("./server.crt", "./server.key"); err != nil {
				if changeHttpsEnableByUser && err == http.ErrServerClosed {
					changeHttpsEnableByUser = false
				} else {
					panic(err)
				}
			}
		} else {
			fmt.Println("run server with https disable...")
			if err := currentServer.ListenAndServe(); err != nil {
				if changeHttpsEnableByUser && err == http.ErrServerClosed {
					changeHttpsEnableByUser = false
				} else {
					panic(err)
				}
			}
		}
	}()

}

func restartServer() {
	fmt.Println("==================================")
	fmt.Println(" server restarting...")
	fmt.Println("==================================")
	shutdownServer()
	runServer()
	fmt.Println("==================================")
	fmt.Println(" server restarting successfully")
	fmt.Println("==================================")
}

func shutdownServer() {
	if currentServer != nil {
		if err := currentServer.Shutdown(context.Background()); err != nil {
			fmt.Printf("shutdown server error: <%v>", err)
			return
		} else {
			fmt.Println("shutdown server successfully")
		}
	}
}

func getIsHttpsEnabled() string {
	if isHttpsEnabled == nil {
		return "nil"
	} else {
		if *isHttpsEnabled {
			return "true"
		} else {
			return "false"
		}
	}
}
