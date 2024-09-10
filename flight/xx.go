package flight

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

func CreateFlightHandler(c *gin.Context) {
	var f Flight
	err := c.ShouldBindJSON(&f)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	f.ID = len(flights) + 1

	flights = append(flights, f)

	c.JSON(http.StatusCreated, f)
}

func GetAllFlightsHandler(c *gin.Context) {
	c.JSON(200, flights)
}

func Handler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
