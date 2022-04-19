package handlers

import (
	t "06-webserver/middlewares"
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World Root")
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the API Endpoint")
}

func PostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user t.User
	err := decoder.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: %s", err)
		return
	}
	name := user.Name
	email := user.Email
	phone := user.Phone
	fmt.Println("Name:, Email:, Phone:", name, email, phone)

	response, err := user.ToJson()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// fmt.Fprintf(w, "Post request: %v\n", user)
	w.Header().Set("Content-Type", "application/json")
	// Write espera un slice de bytes
	w.Write(response)
}
