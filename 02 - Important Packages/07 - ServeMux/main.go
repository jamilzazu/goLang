package main

import "net/http"

func newServerAccount() *http.ServeMux {
	serverAccount := http.NewServeMux()
	serverAccount.HandleFunc("/", handlerAccount)
	serverAccount.Handle("/account", account{name: "Jamil"})
	return serverAccount
}
func handlerAccount(response http.ResponseWriter, _ *http.Request) {
	response.Write([]byte("Server account"))
}

type account struct {
	name string
}

func (account account) ServeHTTP(response http.ResponseWriter, _ *http.Request) {
	response.Write([]byte(account.name))
}

func newServerClient() *http.ServeMux {
	serverClient := http.NewServeMux()
	serverClient.HandleFunc("/", handlerClient)
	return serverClient

}

func handlerClient(response http.ResponseWriter, _ *http.Request) {
	response.Write([]byte("Server client"))
}

func startServers() {
	finish := make(chan bool)

	go func() {
		http.ListenAndServe(":8080", newServerAccount())
	}()

	go func() {
		http.ListenAndServe(":8081", newServerClient())
	}()

	<-finish
}

func main() {
	startServers()
}
