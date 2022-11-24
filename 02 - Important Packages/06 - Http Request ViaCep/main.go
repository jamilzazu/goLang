package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

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

func searchZipCode(zipCode string) (*ViaCep, error) {
	url := "https://viacep.com.br/ws/" + zipCode + "/json/"
	response, error := http.Get(url)
	messageErrorViaCep(error)

	defer response.Body.Close()

	body, error := ioutil.ReadAll(response.Body)
	messageErrorViaCep(error)

	var code ViaCep
	error = json.Unmarshal(body, &code)

	messageErrorViaCep(error)
	return &code, nil
}

func messageErrorViaCep(err error) (*ViaCep, error) {
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func searchZipCodeHandler(response http.ResponseWriter, request *http.Request) {
	url := request.URL
	zipCode := url.Query().Get("cep")
	UrlIsDifferentFromSlash(response, url.Path)
	zipCodeEmpty(response, zipCode)

	result, error := searchZipCode(zipCode)
	zipCodeRequestError(response, error)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}

func UrlIsDifferentFromSlash(response http.ResponseWriter, urlPath string) {
	if urlPath != "/" {
		response.WriteHeader(http.StatusNotFound)
	}
}

func zipCodeEmpty(response http.ResponseWriter, zipCode string) {
	if zipCode == "" {
		response.WriteHeader(http.StatusBadRequest)
	}
}

func zipCodeRequestError(response http.ResponseWriter, err error) {
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", searchZipCodeHandler)
	http.ListenAndServe(":8080", nil)
}
