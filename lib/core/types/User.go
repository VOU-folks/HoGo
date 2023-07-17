package types

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Roles    []Role `json:"roles"`
	Active   bool   `json:"active"`
	Deleted  bool   `json:"deleted"`
}
