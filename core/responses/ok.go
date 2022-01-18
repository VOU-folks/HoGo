package responses

import (
	"net/http"

	"hogo/core/components"
)

func OK(c *components.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}
