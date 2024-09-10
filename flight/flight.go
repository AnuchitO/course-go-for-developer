package flight

import (
	"github.com/gin-gonic/gin"
)

func New() *Flight {
	return &Flight{}
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

func Handler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
