package responses

import (
	"fmt"
	"net/http"

	"github.com/go-errors/errors"

	"hogo/core/components"
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

func HandleSystemError(c *components.Context, err error) {
	c.AbortWithStatusJSON(
		http.StatusInternalServerError,
		NewSystemError(err))
}

func RespondWithSystemError(c *components.Context, err error) {
	HandleSystemError(c, err)
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
			RespondWithSystemError(c, err)
		}
	}(c)
	c.Next()
}
