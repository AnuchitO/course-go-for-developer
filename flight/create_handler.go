package flight

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateHandler(c *gin.Context) {
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
