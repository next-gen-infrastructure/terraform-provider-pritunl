package pritunl

type LocationRoute struct {
	ID         string `json:"id,omitempty"`
	Network    string `json:"network"`
	LinkId     string `json:"link_id"`
	LocationId string `json:"location_id"`
}
