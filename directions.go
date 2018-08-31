package mapfit

import (
	"bytes"
	"encoding/json"
	"fmt"
)

var (
	directionsApi = &apiConfig{
		method:  "POST",
		baseUrl: "https://api.mapfit.com",
		path:    "/v2/directions",
	}

	AvailableDirectionsType = []string{"driving", "walking"}
)

// DirectionsResponse
// Thanks https://mholt.github.io/json-to-go/ ;)
type DirectionsResponse struct {
	Trip struct {
		Summary struct {
			MinLon float64 `json:"min_lon"`
			MaxLat float64 `json:"max_lat"`
			MaxLon float64 `json:"max_lon"`
			Length float64 `json:"length"`
			Time   int     `json:"time"`
			MinLat float64 `json:"min_lat"`
		} `json:"summary"`
		StatusMessage string `json:"status_message"`
		Legs          []struct {
			Summary struct {
				MinLon float64 `json:"min_lon"`
				MaxLat float64 `json:"max_lat"`
				MaxLon float64 `json:"max_lon"`
				Length float64 `json:"length"`
				Time   int     `json:"time"`
				MinLat float64 `json:"min_lat"`
			} `json:"summary"`
			Shape     string `json:"shape"`
			Maneuvers []struct {
				VerbalMultiCue                   bool     `json:"verbal_multi_cue,omitempty"`
				BeginShapeIndex                  int      `json:"begin_shape_index"`
				TravelMode                       string   `json:"travel_mode"`
				Instruction                      string   `json:"instruction"`
				StreetNames                      []string `json:"street_names,omitempty"`
				Length                           float64  `json:"length"`
				EndShapeIndex                    int      `json:"end_shape_index"`
				Time                             int      `json:"time"`
				Type                             int      `json:"type"`
				VerbalPreTransitionInstruction   string   `json:"verbal_pre_transition_instruction"`
				TravelType                       string   `json:"travel_type"`
				VerbalTransitionAlertInstruction string   `json:"verbal_transition_alert_instruction,omitempty"`
				VerbalPostTransitionInstruction  string   `json:"verbal_post_transition_instruction,omitempty"`
				Sign                             struct {
					ExitTowardElements []struct {
						Text string `json:"text"`
					} `json:"exit_toward_elements"`
				} `json:"sign,omitempty"`
				Toll             bool     `json:"toll,omitempty"`
				BeginStreetNames []string `json:"begin_street_names,omitempty"`
			} `json:"maneuvers"`
		} `json:"legs"`
		Language  string `json:"language"`
		Locations []struct {
			Lon          float64 `json:"lon"`
			SideOfStreet string  `json:"side_of_street"`
			Type         string  `json:"type"`
			Lat          float64 `json:"lat"`
		} `json:"locations"`
		Units  string `json:"units"`
		Status int    `json:"status"`
	} `json:"trip"`
	DestinationLocation []float64 `json:"destinationLocation"`
	SourceLocation      []float64 `json:"sourceLocation"`
}

// GetDirectionsWithAddress
func (c *Client) GetDirections(srcAddr, dstAddr Address, dType string) (*DirectionsResponse, error) {
	if isDirectionTypeAvailable(dType) == false {
		return nil, fmt.Errorf("Directions type %s don't exist", dType)
	}

	// Json structure for direction request
	directionsRequest := struct {
		SrcAddr Address `json:"source-address"`
		DstAddr Address `json:"destination-address"`
		Type    string  `json:"type"`
	}{
		SrcAddr: srcAddr,
		DstAddr: dstAddr,
		Type:    dType,
	}
	jsonReq, err := json.Marshal(directionsRequest)
	if err != nil {
		return nil, err
	}
	body := bytes.NewBuffer(jsonReq)

	resp, err := c.request(directionsApi, body)
	if err != nil {
		return nil, err
	}

	directionsResp := &DirectionsResponse{}
	err = json.Unmarshal(resp, directionsResp)
	return directionsResp, err
}

func isDirectionTypeAvailable(dType string) bool {
	for _, t := range AvailableDirectionsType {
		if t == dType {
			return true
		}
	}
	return false
}
