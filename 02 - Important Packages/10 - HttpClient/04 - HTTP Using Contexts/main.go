package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"
)

func messageError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	newContext := context.Background()
	newContext, cancel := context.WithTimeout(newContext, time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(newContext, "GET", "http://google.com", nil)
	messageError(err)

	resp, err := http.DefaultClient.Do(req)
	messageError(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	messageError(err)
	println(string(body))
}
