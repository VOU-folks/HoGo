package responses

import (
  "net/http"

  "hogo/core/components"
)

type BadRequest struct {
  Type    string `json:"type"`
  Message string `json:"message"`
}

func NewBadRequest() BadRequest {
  return BadRequest{
    Type:    "bad-request",
    Message: "Request requirements does not match",
  }
}

func NewBadRequestWithMessage(message string) BadRequest {
  return BadRequest{
    Type:    "bad-request",
    Message: message,
  }
}

func RespondWithBadRequest(c *components.Context) {
  c.AbortWithStatusJSON(
    http.StatusBadRequest,
    NewBadRequest())
}

func RespondWithMessageInBadRequest(message string, c *components.Context) {
  c.AbortWithStatusJSON(
    http.StatusBadRequest,
    NewBadRequestWithMessage(message))
}

func HandleBadRequest(c *components.Context) {
  RespondWithBadRequest(c)
}
