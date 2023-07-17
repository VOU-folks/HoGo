package types

import "time"

type Instance struct {
	Id       string       `json:"id"`
	Name     string       `json:"name"`
	Role     InstanceRole `json:"role"`
	IP       string       `json:"ip"`
	Hostname string       `json:"hostname"`
	LastSync time.Time    `json:"lastSync"`
	Status   string       `json:"status"`
	Active   bool         `json:"active"`
	Deleted  bool         `json:"deleted"`
}

type InstanceRole string

const (
	InstanceRoleNameserver  = "nameserver"
	InstanceRoleApplication = "application"
	InstanceRoleStorage     = "storage"
	InstanceRoleDatabase    = "database"
)
