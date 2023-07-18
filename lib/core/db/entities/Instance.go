package entities

import (
	"hogo/lib/core/db/enums"
)

type Instance struct {
	Entity

	Id       string             `json:"id"`
	Name     string             `json:"name"`
	Role     enums.InstanceRole `json:"role"`
	IP       string             `json:"ip"`
	Hostname string             `json:"hostname"`
	LastSync NullTime           `json:"lastSync"`
	Status   string             `json:"status"`
	Active   bool               `json:"active"`
	Deleted  bool               `json:"deleted"`
}
