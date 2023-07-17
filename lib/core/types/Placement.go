package types

type Placement struct {
	Id        string     `json:"id"`
	Instances []Instance `json:"instance"`
	Locked    bool       `json:"locked"`
	Active    bool       `json:"active"`
	Deleted   bool       `json:"deleted"`
}
