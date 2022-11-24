package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// convert json to struct: https://mholt.github.io/json-to-go/

type ViaCep struct {
	ZipCode      string `json:"cep"`
	Street       string `json:"logradouro"`
	Complement   string `json:"complemento"`
	Neighborhood string `json:"bairro"`
	City         string `json:"localidade"`
	State        string `json:"uf"`
	Ibge         string `json:"ibge"`
	Gia          string `json:"gia"`
	Ddd          string `json:"ddd"`
	Siafi        string `json:"siafi"`
}

func textInput() {
	for _, zipcode := range os.Args[1:] {
		url := "https://viacep.com.br/ws/" + zipcode + "/json/"
		request := searchZipCode(url)
		response := getResponse(request)
		jsonUnmarshal(response)
	}
}

func searchZipCode(url string) *http.Response {
	request, err := http.Get(url)
	messageErrorRequest(err)
	return request
}

func getResponse(request *http.Response) []byte {
	response, err := io.ReadAll(request.Body)
	messageErrorResponse(err)
	closeRequestBody(request)
	return response
}

func closeRequestBody(request *http.Response) {
	defer request.Body.Close()
}

func jsonUnmarshal(response []byte) {
	var data ViaCep
	var err = json.Unmarshal(response, &data)
	messageErrorResponseParse(err)
	fmt.Println(data)
	createFile(data)
}

func createFile(data ViaCep) {
	file, err := os.Create("city.txt")
	messageErrorCreateFile(err)
	_, err = file.WriteString(fmt.Sprintf("ZipCode: %s, City: %s, State: %s", data.ZipCode, data.City, data.State))
	closeFile(file)
}

func closeFile(file *os.File) {
	defer file.Close()
}

func messageErrorRequest(message error) {
	if message != nil {
		fmt.Fprintf(os.Stderr, "Error when making request: %v \n", message)
	}
}

func messageErrorResponse(message error) {
	if message != nil {
		fmt.Fprintf(os.Stderr, "Error reading response: %v \n", message)
	}
}

func messageErrorResponseParse(message error) {
	if message != nil {
		fmt.Fprintf(os.Stderr, "Error parsing response: %v \n", message)
	}
}

func messageErrorCreateFile(message error) {
	if message != nil {
		fmt.Fprintf(os.Stderr, "Error create file: %v \n", message)
	}
}

func main() {
	textInput()
}
