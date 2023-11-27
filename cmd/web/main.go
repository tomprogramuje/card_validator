package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tomprogramuje/card_validator/validation"
)

const port = ":80"

type jsonRequest struct {
	CreditCardNumber string `json:"creditCardNumber"`
}

type jsonResponse struct {
	Valid bool `json:"valid"`
	Issuer string `json:"issuer"`
}

func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	// checks if the http method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// decodes json payload 
	var reqBody jsonRequest
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// validates the card number
	valid, issuer := validation.CheckCardNumber(reqBody.CreditCardNumber)

	// makes json response
	resp := jsonResponse{Valid: valid, Issuer: issuer}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/", handlePostRequest)

	log.Println("Starting app on port", port)
	err := http.ListenAndServe(port, nil)
	log.Println(err)
}
