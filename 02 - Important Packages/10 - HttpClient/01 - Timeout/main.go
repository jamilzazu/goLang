package main

import (
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
	c := http.Client{Timeout: time.Second}
	resp, err := c.Get("http://google.com")
	messageError(err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	messageError(err)

	println(string(body))
}
