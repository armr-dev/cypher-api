package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/armr-dev/cypher-api-go/pkg/cmd/3des"
	"github.com/armr-dev/cypher-api-go/pkg/cmd/blowfish"
	"github.com/armr-dev/cypher-api-go/pkg/cmd/des"
	"github.com/gorilla/mux"
)

type Request struct {
	Text      string `json:"text"`
	Algorithm string `json:"algorithm"`
}

type Response struct {
	Text string `json:"text"`
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

	switch request.Algorithm {
	case "des":
		encryptedText, _ = DES.Encrypt(request.Text)

	case "3des":
		encryptedText, _ = TripleDES.Encrypt(request.Text)

	case "blowfish":
	default:
		encryptedText = Blowfish.Encrypt(request.Text)
	}

	var response Response

	response.Text = encryptedText

	json.NewEncoder(w).Encode(response)
}

func decipherText(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var request Request

	json.Unmarshal(reqBody, &request)

	var decryptedText string

	switch request.Algorithm {
	case "des":
		decryptedText, _ = DES.Decrypt(request.Text)

	case "3des":
		decryptedText, _ = TripleDES.Decrypt(request.Text)

	case "blowfish":
	default:
		decryptedText, _ = Blowfish.Decrypt(request.Text)
	}

	var response Response

	response.Text = decryptedText

	json.NewEncoder(w).Encode(response)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/cypher", cypherText).Methods("POST")
	myRouter.HandleFunc("/decipher", decipherText).Methods("POST")

	var port string

	_, envExists := os.LookupEnv("PORT")

	if envExists {
		port = os.Getenv("PORT")
	} else {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(":"+port, myRouter))
}

func main() {
	handleRequests()
}
