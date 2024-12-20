package pritunl

type Location struct {
	ID       string          `json:"id,omitempty"`
	Name     string          `json:"name"`
	Type     string          `json:"type,omitempty"`
	IpV6     bool            `json:"ipv6,omitempty"`
	LinkId   string          `json:"link_id"`
	LinkType string          `json:"link_type,omitempty"`
	Hosts    []LocationHost  `json:"hosts,omitempty"`
	Routes   []LocationRoute `json:"routes,omitempty"`
	Peers    []any           `json:"peers,omitempty"`
}
