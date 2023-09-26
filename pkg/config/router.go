package config

type RouterSpecifies struct {
	Mss     int      `json:"tcpMss,omitempty"`
	Name    string   `json:"name,omitempty"`
	Link    string   `json:"link,omitempty"`
	Subnets []Subnet `json:"subnets"`
}

func (n *RouterSpecifies) Correct() {
}
