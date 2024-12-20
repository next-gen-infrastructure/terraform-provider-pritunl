package pritunl

type URI struct {
	ID                  string `json:"id"`
	Name                string `json:"name"`
	LinkId              string `json:"link_id"`
	LocationId          string `json:"location_id"`
	Status              string `json:"status"`
	Hosts               string `json:"hosts,omitempty"`
	HostsStateAvailable int    `json:"hosts_state_available"`
	HostsStateTotal     int    `json:"hosts_state_total"`
	Timeout             int    `json:"timeout,omitempty"`
	Priority            int    `json:"priority"`
	Backoff             int    `json:"backoff,omitempty"`
	PingTimestampTTL    int    `json:"ping_timestamp_ttl,omitempty"`
	Static              bool   `json:"static"`
	PublicAddress       string `json:"public_address"`
	LocalAddress        string `json:"local_address"`
	Address6            string `json:"address6,omitempty"`
	Version             string `json:"version,omitempty"`
	URI                 string `json:"uri"`
}
