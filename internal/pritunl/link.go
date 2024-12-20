package pritunl

type Link struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Type           string `json:"type"`
	Action         string `json:"action"`
	PreferredIKE   string `json:"preferred_ike"`
	PreferredESP   string `json:"preferred_esp"`
	HostCheck      bool   `json:"host_check"`
	IPv6           bool   `json:"ipv6"`
	ForcePreferred bool   `json:"force_preferred"`
}

type Links struct {
	Page      int    `json:"page"`
	PageTotal int    `json:"page_total"`
	Links     []Link `json:"links"`
}