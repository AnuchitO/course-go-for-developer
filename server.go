package main

import (
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

func main() {
	r := gin.Default()

	r.GET("/ping", handler)
	// GET /flights/1 ->  GetFlightByIDHandler ->
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
