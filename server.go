package main

import (
	"net/http"
	"strconv"

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

func handler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

type Flight struct {
	ID          int
	Number      int
	AirlineCode string
	Destination string
	Arrival     string
}

var flights = []Flight{
	{ID: 1, Number: 3250, AirlineCode: "FD", Destination: "DMK", Arrival: "KKC"},
	{ID: 2, Number: 3251, AirlineCode: "FD", Destination: "DMK", Arrival: "KKC"},
}

func GetFlightByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	// l := len(flights)
	// for i := 0; i < l; i++ {
	// 	f := flights[i]
	// 	if f.ID == id {
	// 		c.JSON(200, f)
	// 		return
	// 	}
	// }
	// for i := range flights {
	// 	f := flights[i]
	// 	if f.ID == id {
	// 		c.JSON(200, f)
	// 		return
	// 	}
	// }

	for _, f := range flights {
		if f.ID == id {
			c.JSON(200, f)
			return
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "flight ID not found "})
}

func main() {
	r := gin.Default()

	r.GET("/ping", handler)
	r.GET("/flights/:id", GetFlightByIDHandler)
	// POST /flights => CreateFlightHandler playload { "number": 123, "airlineCode": "AA", "destination": "DMK", "arrival": "KKC" }
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
