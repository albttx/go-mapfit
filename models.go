package mapfit

// Address direction request
type Address struct {
	StreetAddress string `json:"street_address,omitempty"`
	Locality      string `json:"locality,omitempty"`
	Admin         string `json:"admin_1,omitempty"`
	Zip           string `json:"zip,omitempty"`
	EntranceType  string `json:"entrance_type,omitempty"`

	Lat string `json:"lat,omitempty"`
	Lon string `json:"lon,omitempty"`
}
