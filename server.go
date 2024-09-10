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

func GetFlightByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	f := Flight{
		ID:          id,
		Number:      3250,
		AirlineCode: "FD",
		Destination: "DMK",
		Arrival:     "KKC",
	}

	c.JSON(200, f)
}

func main() {
	r := gin.Default()

	r.GET("/ping", handler)
	r.GET("/flights/:id", GetFlightByIDHandler)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
