package types

type Project struct {
	Id            string    `json:"id"`
	Name          string    `json:"name"`
	Owner         User      `json:"owner"`
	Collaborators []User    `json:"collaborators"`
	Placement     Placement `json:"placement"`
	DNSZones      []DNSZone `json:"dnsZones"`
	Active        bool      `json:"active"`
	Deleted       bool      `json:"deleted"`
}
