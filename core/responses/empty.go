package responses

import (
  "net/http"

  "hogo/core/components"
)

type Empty struct{}

func RespondWithEmpty(c *components.Context) {
  c.Status(http.StatusNoContent)
}

func HandleEmpty(c *components.Context) {
  RespondWithEmpty(c)
}
