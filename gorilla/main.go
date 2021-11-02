package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type API struct {
	Message string `json:"message"`
}

func hello(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	id := urlParams["id"]
	messageContact := "Kullanıcı ID: " + id

	message := API{messageContact}
	output, err := json.Marshal(message)

	if err != nil {
		fmt.Println("error oluştu")
	}
	fmt.Fprintf(w, string(output))

}

func main() {
	gorillaRoute := mux.NewRouter()
	gorillaRoute.HandleFunc("/api/user{id}", hello)
	http.Handle("/", gorillaRoute)
	http.ListenAndServe(":3000", nil)
}
