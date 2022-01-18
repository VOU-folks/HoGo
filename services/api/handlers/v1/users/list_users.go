package users

import (
	"github.com/jkomyno/nanoid"

	"hogo/core/components"
	"hogo/core/responses"
)

type User struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type UserListing struct {
	Users []User `json:"users" default:"[]"`
}

func ListUsers(c *components.Context) {
	var users = make([]User, 0)

	id, _ := nanoid.Nanoid(16)
	users = append(users, User{
		Id:        id,
		FirstName: "Anar",
		LastName:  "Jafarov",
	})

	responses.OK(c, UserListing{Users: users})
}
