package pritunl

type LocationHost struct {
	ID                  string  `json:"id"`
	Name                string  `json:"name"`
	LinkID              string  `json:"link_id"`
	LocationID          string  `json:"location_id"`
	Status              string  `json:"status"`
	Hosts               *string `json:"hosts"` // Use *string to allow null values
	HostsStateAvailable int     `json:"hosts_state_available"`
	HostsStateTotal     int     `json:"hosts_state_total"`
	Timeout             *int    `json:"timeout"` // Use *int to allow null values
	Priority            int     `json:"priority"`
	Backoff             *int    `json:"backoff"`            // Use *int to allow null values
	PingTimestampTTL    *string `json:"ping_timestamp_ttl"` // Use *int to allow null values
	Static              bool    `json:"static"`
	PublicAddress       string  `json:"public_address"`
	LocalAddress        string  `json:"local_address"`
	Address6            *string `json:"address6"` // Use *string to allow null values
	Version             *string `json:"version"`  // Use *string to allow null values
	URI                 string  `json:"uri"`
}
