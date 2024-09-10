package main

import (
	"encoding/json"
	"log"
)

type Flight struct {
	ID          int    `json:"id"`
	Number      int    `json:"number"`
	AirlineCode string `json:"airline_code"`
	Destination string `json:"destination"`
	Arrival     string `json:"arrival"`
}

func main() {
	// using encoding/json
	// struct -> JSON ([]byte of text)   ==> Marshal (Serialize)
	// JSON ([]byte of text) -> struct 	 ==> Unmarshal (Deserialize)

	f := Flight{ID: 1, Number: 3250, AirlineCode: "FD", Destination: "DMK", Arrival: "KKC"}

	log.Printf("%#v\n", f)

	b, err := json.Marshal(f)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(b))
}
