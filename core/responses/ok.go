package responses

import (
	"net/http"

	"hogo/core/components"
)

func OK(data interface{}, c *components.Context) {
	c.JSON(http.StatusOK, data)
}
