package ginadapter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/italorfeitosa/go-ports-n-adapters/internal/port/primary"
)

func Adapt[In any, Req RequestMapper[In], Out any, Res ResponseMapper[Out]](
	requestMapper Req,
	featureFn primary.FeatureFunc[In, Out],
	responseMapper Res,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestMapper Req
		var responseMapper Res

		if err := c.ShouldBindUri(&requestMapper); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := c.ShouldBindQuery(&requestMapper); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := c.ShouldBindJSON(&requestMapper); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		result, err := featureFn(requestMapper.ToInput())

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		if responseMapper.StatusCode() == http.StatusNoContent {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.JSON(responseMapper.StatusCode(), responseMapper.FromOutput(result))
	}
}
