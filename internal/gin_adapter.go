package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdaptController[In any, Out any](successStatusCode int, execController ControllerFunc[In, Out]) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req In

		if err := c.ShouldBindUri(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := c.ShouldBindQuery(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		result, err := execController(req)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(successStatusCode, result)
	}
}
