package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/flights", func(w http.ResponseWriter, r *http.Request) {
		raw := `{
		"name": "anuchitO",
		"age": 25
		}`
		w.Write([]byte(raw))
	})

	log.Println("start server at port :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
