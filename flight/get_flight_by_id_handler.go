package flight

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetByIDHandler(c *gin.Context) {
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
