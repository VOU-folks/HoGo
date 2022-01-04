package responses

import (
  "net/http"

  "github.com/go-playground/validator/v10"
  "hogo/core/components"
)

type InvalidField struct {
  Field   string `json:"field"`
  Message string `json:"message"`
}

type InvalidRequest struct {
  Type    string         `json:"type"`
  Message string         `json:"message"`
  Fields  []InvalidField `json:"fields"`
}

func NewInvalidRequest(invalidFields []InvalidField) InvalidRequest {
  return InvalidRequest{
    Type:    "validation",
    Message: "One or more field(s) not valid",
    Fields:  invalidFields,
  }
}

func RespondWithInvalidRequest(err error, c *components.Context) {
  switch err.(type) {
  case *validator.ValidationErrors:
    var invalidFields []InvalidField
    for _, invalidField := range err.(validator.ValidationErrors) {
      invalidFields = append(
        invalidFields,
        InvalidField{
          Field:   invalidField.Field(),
          Message: invalidField.Translate(nil),
        })
    }

    c.AbortWithStatusJSON(
      http.StatusBadRequest,
      NewInvalidRequest(invalidFields))

  default:
    RespondWithMessageInBadRequest(err.Error(), c)
  }
}
