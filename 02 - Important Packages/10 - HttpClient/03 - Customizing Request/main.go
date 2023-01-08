package main

import (
	"io/ioutil"
	"net/http"
)

func messageError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	client := http.Client{}
	req, err := http.NewRequest("GET", "http://google.com", nil)
	messageError(err)
	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)
	messageError(err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	messageError(err)
	println(string(body))
}
