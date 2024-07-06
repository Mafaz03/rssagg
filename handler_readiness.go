package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handler_readiness(w http.ResponseWriter, r *http.Request) {

	type parameter struct {
		Message string `json:"Satus code [200]"`
	}
	params := parameter{
		Message: "You are ready to go!",
	}

	jsonData, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err)
		return
	}

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&jsonData)
	if err != nil {
		log.Printf("There was an error decoding the params")
	}

	respondWithJSON(w, 200, params)
}
