package teams

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// swagger:parameters getTeams
type HandlerRequest struct {
}

// swagger:response getTeams
type HandlerResponse struct {
	// in:body
	Teams
}

func Handler(
	c *gin.Context,
) {

	var err error

	response := HandlerResponse{}

	if response.Teams, err = Get(c.Request.Context(), nil); err != nil {

		c.AbortWithError(http.StatusServiceUnavailable, err)

		return
	}

	c.JSON(http.StatusOK, response.Teams)
}
