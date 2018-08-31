package mapfit

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type ErrorMsg struct {
	code int
	exit bool
	msg  error
}

var (
	geocodeApi = &apiConfig{
		method:  "POST",
		baseUrl: "https://api.mapfit.com",
		path:    "/v2/geocode",
	}

	errorsMsg = []ErrorMsg{
		{code: 0, exit: true, msg: errors.New("Must specify a value for 'street_address'")},
		{code: 2, exit: false, msg: errors.New("Interpolated address. This means the exact address was not found but the information given was enough for an educated guess.")},
		{code: 3, exit: true, msg: errors.New("Zero results were found from the input query")},
		{code: 5, exit: false, msg: errors.New("Region. We found a geographic region by the name provided in the street_address field. This can either be a locality, an admin_1, or a country. The response street_address will contain a label suitable for display of the region on a map. For example, a result for a search with street_address 'Washington DC’ is street_address 'Washington, District Of Columbia, United States'.")},
		{code: 6, exit: false, msg: errors.New("Road. We found a road the provided the street_address as its name")},
		{code: 13, exit: false, msg: errors.New("Fallback. We relied on a less-precise geocoding approach to find the input. The result is either a building centroid, best-guess, geographic region or other data type about which we make very few guarantees.")},
	}
)

type GeocodeResponse struct {
	Locality     string `json:"locality"`
	PostalCode   string `json:"postal_code"`
	Admin1       string `json:"admin_1"`
	Country      string `json:"country"`
	Neighborhood string `json:"neighborhood"`
	ResponseType int    `json:"response_type"`
	Message      string `json:"message"`
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
	if err = json.Unmarshal(resp, &geocodeResp); err != nil {
		return geocodeResp, err
	}
	if geocodeResp[0].ResponseType != 1 {
		code := geocodeResp[0].ResponseType
		for _, e := range errorsMsg {
			if code == e.code && e.exit {
				return geocodeResp, fmt.Errorf("code=%v, message=%s", code, e.msg)
			}
		}
	}
	return geocodeResp, nil
}
