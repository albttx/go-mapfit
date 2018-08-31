package mapfit

import (
	"bytes"
	"encoding/json"
)

var (
	geocodeApi = &apiConfig{
		method:  "POST",
		baseUrl: "https://api.mapfit.com",
		path:    "/v2/geocode",
	}
)

type GeocodeResponse struct {
	Locality     string `json:"locality"`
	PostalCode   string `json:"postal_code"`
	Admin1       string `json:"admin_1"`
	Country      string `json:"country"`
	Neighborhood string `json:"neighborhood"`
	ResponseType int    `json:"response_type"`
	Building     struct {
		Type        string        `json:"type"`
		Coordinates [][][]float64 `json:"coordinates"`
	} `json:"building"`
	StreetAddress string `json:"street_address"`
	Entrances     []struct {
		Lon          float64 `json:"lon"`
		Lat          float64 `json:"lat"`
		EntranceType string  `json:"entrance_type"`
	} `json:"entrances"`
}

// Geocode
func (c *Client) Geocode(addr Address) ([]GeocodeResponse, error) {
	jsonReq, err := json.Marshal(addr)
	if err != nil {
		return nil, err
	}
	body := bytes.NewBuffer(jsonReq)

	resp, err := c.request(geocodeApi, body)
	if err != nil {
		return nil, err
	}
	geocodeResp := []GeocodeResponse{}
	err = json.Unmarshal(resp, &geocodeResp)
	return geocodeResp, err
}
