package main

import (
	"github.com/anuchito/demoapi/flight"
	"github.com/gin-gonic/gin"
)

/*
GET /flights/{id}

Response:

	{
		"id": 1,
		"airlineCode": "FD",
		"number": "3250",
		"departure": "DMK",
		"arrival": "KKC",
	}
*/

func main() {
	r := gin.Default()

	r.GET("/ping", flight.Handler)
	r.GET("/flights", flight.GetAllFlightsHandler)     // SELECT * FROM flights;
	r.GET("/flights/:id", flight.GetFlightByIDHandler) // SELECT * FROM flights WHERE id = ?;
	r.POST("/flights", flight.CreateFlightHandler)     // INSERT INTO flights (id, number, airlineCode, destination, arrival) VALUES (?, ?, ?, ?, ?);
	// POST /flights
	// POST /flights/:id

	// POST /flights/actions/create
	// POST /flights/actions/update
	// POST /flights/actions/delete

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
