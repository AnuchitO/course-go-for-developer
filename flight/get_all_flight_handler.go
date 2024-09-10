package flight

import "github.com/gin-gonic/gin"

func GetAllHandler(c *gin.Context) {
	c.JSON(200, flights)
}
