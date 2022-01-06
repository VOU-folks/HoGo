package responses

import (
	"net/http"

	"hogo/core/components"
)

type NotFound struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func NewNotFound() NotFound {
	return NotFound{
		Type:    "not-found",
		Message: "Requested network resource does not exist",
	}
}

func HandleNotFound(c *components.Context) {
	c.AbortWithStatusJSON(
		http.StatusNotFound,
		NewNotFound())
}

func RespondWithNotFound(c *components.Context) {
	HandleNotFound(c)
}
