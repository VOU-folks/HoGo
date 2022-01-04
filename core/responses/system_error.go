package responses

import (
  "fmt"
  "net/http"

  "hogo/core/components"

  "github.com/go-errors/errors"
)

type SystemError struct {
  Type    string `json:"type"`
  Message string `json:"message"`
}

func NewSystemError(err error) SystemError {
  return SystemError{
    Type:    "system",
    Message: err.Error(),
  }
}

func HandleSystemError(err error, c *components.Context) {
  c.AbortWithStatusJSON(
    http.StatusInternalServerError,
    NewSystemError(err))
}

func RespondWithSystemError(err error, c *components.Context) {
  HandleSystemError(err, c)
}

func RecoverOnSystemError(c *components.Context) {
  defer func(c *components.Context) {
    if rec := recover(); rec != nil {
      var err error
      switch t := rec.(type) {
      case string:
        err = errors.New(t)
      case error:
        err = t
      default:
        err = errors.New("Unknown error")
      }
      fmt.Println(err)
      RespondWithSystemError(err, c)
    }
  }(c)
  c.Next()
}
