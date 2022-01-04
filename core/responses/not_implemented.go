package responses

import (
  "net/http"

  "hogo/core/components"
)

type NotImplemented struct {
  Type    string `json:"type"`
  Message string `json:"message"`
}

func NewNotImplemented() NotImplemented {
  return NotImplemented{
    Type:    "not-implemented",
    Message: "Requested method not implemented",
  }
}

func HandleNotImplemented(c *components.Context) {
  c.AbortWithStatusJSON(
    http.StatusNotImplemented,
    NewNotImplemented())
}

func RespondWithNotImplemented(c *components.Context) {
  HandleNotImplemented(c)
}
