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
	"github.com/armr-dev/cypher-api-go/pkg/cmd/aes"
	"github.com/gorilla/mux"
)

type Request struct {
	Text      string `json:"text"`
	Algorithm string `json:"algorithm"`
}

type Data struct {
	Text string `json:"text"`
}

type Response struct {
	Data `json:"data"`
}

func setupResponse(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	fmt.Fprintf(w, "Welcome to the home page!")
	fmt.Println("Endpoint Hit: homepage")
}

func cypherText(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	reqBody, _ := ioutil.ReadAll(r.Body)
	
	var request Request
	
	json.Unmarshal(reqBody, &request)
	
	var encryptedText string
	
	switch request.Algorithm {
	case "aes":
		encryptedText, _ = AES.Encrypt(request.Text)

	case "des":
		encryptedText, _ = DES.Encrypt(request.Text)
		
	case "3des":
		encryptedText, _ = TripleDES.Encrypt(request.Text)
		
	case "blowfish":
		encryptedText = Blowfish.Encrypt(request.Text)
	
	default:
		encryptedText = "Error"
	}
	
	var response Response
	
	response.Data.Text = encryptedText
	
	json.NewEncoder(w).Encode(response)
}

func decipherText(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	reqBody, _ := ioutil.ReadAll(r.Body)

	var request Request

	json.Unmarshal(reqBody, &request)

	var decryptedText string

	switch request.Algorithm {
	case "aes":
		decryptedText, _ = AES.Decrypt(request.Text)

	case "des":
		decryptedText, _ = DES.Decrypt(request.Text)

	case "3des":
		decryptedText, _ = TripleDES.Decrypt(request.Text)

	case "blowfish":
		decryptedText, _ = Blowfish.Decrypt(request.Text)

	default:
		decryptedText = "Error"
	}

	var response Response

	response.Data.Text = decryptedText

	json.NewEncoder(w).Encode(response)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/cypher", cypherText).Methods("POST", "OPTIONS")
	myRouter.HandleFunc("/decipher", decipherText).Methods("POST", "OPTIONS")

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
