package players

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// swagger:parameters getPlayers
type HandlerRequest struct {
	TeamID *uuid.UUID `json:"teamId"`
}

// swagger:response getPlayers
type HandlerResponse struct {
	// in:body
	Players
}

func Handler(
	c *gin.Context,
) {

	var err error

	response := HandlerResponse{}

	request := HandlerRequest{}

	if response.Players, err = Get(c.Request.Context(), request.TeamID); err != nil {

		c.AbortWithError(http.StatusServiceUnavailable, err)

		return
	}

	c.JSON(http.StatusOK, response.Players)
}
