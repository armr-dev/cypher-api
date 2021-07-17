package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/armr-dev/cypher-api-go/pkg/cmd/des"
	"github.com/armr-dev/cypher-api-go/pkg/cmd/3des"
	"github.com/armr-dev/cypher-api-go/pkg/cmd/blowfish"
	"github.com/gorilla/mux"
)

type Request struct {
	Text 			string `json:"text"`
	Algorithm string `json:"algorithm"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
	fmt.Println("Endpoint Hit: homepage")
}

func cypherText(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var request Request

	json.Unmarshal(reqBody, &request)

	var encryptedText string

	switch (request.Algorithm) {
	case "des":
		encryptedText, _ = DES.Encrypt(request.Text)

	case "3des":
		encryptedText, _ = TripleDES.Encrypt(request.Text)

	case "blowfish": default:
		encryptedText = Blowfish.Encrypt(request.Text)
	}

	fmt.Fprintf(w, "%+v", encryptedText)
}

func decipherText(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var request Request

	json.Unmarshal(reqBody, &request)

	decryptedText, _ := DES.Decrypt(request.Text)

	fmt.Fprintf(w, "%+v", string(decryptedText))
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/cypher", cypherText)
	myRouter.HandleFunc("/decipher", decipherText)

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main () {
	handleRequests()
}