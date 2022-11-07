package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	resp, err := http.Post("http://10.186.60.99:10000/v1/login", "application/x-www-form-urlencoded", strings.NewReader("username=admin&password=admin"))
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response: [%v]", string(body))
}
