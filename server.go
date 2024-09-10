package main

import (
	"fmt"
	"log"
	"net/http"
)

func flightHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("flightHandler Method: ", r.Method)

	raw := `{
	"name": "anuchitO",
	"age": 25
	}`
	w.Write([]byte(raw))
}

func main() {
	http.HandleFunc("/flights", flightHandler)

	log.Println("start server at port :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
