package entities

import "hogo/lib/core/db/enums"

type User struct {
	Entity

	Id       string           `json:"id"`
	Name     string           `json:"name"`
	Username string           `json:"username"`
	Email    string           `json:"email"`
	Password string           `json:"password"`
	Roles    []enums.UserRole `json:"roles"`
	Active   bool             `json:"active"`
	Deleted  bool             `json:"deleted"`
}
