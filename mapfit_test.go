package mapfit_test

import (
	"fmt"
	"os"
	"testing"

	mapfit "github.com/albttx/go-mapfit"
)

var (
	MapfitToken = os.Getenv("MAPFIT_TOKEN")
)

func init() {
	if MapfitToken == "" {
		fmt.Println("Environment variable missing: MAPFIT_TOKEN")
		os.Exit(1)
	}
}

func TestMapfitDirections(t *testing.T) {
	mapfitClient := mapfit.NewClient(MapfitToken)
	srcAddr := mapfit.Address{
		StreetAddress: "avenue des Champs Elysee",
		Locality:      "Paris",
	}
	dstAddr := mapfit.Address{
		StreetAddress: "Tour Eiffel",
		Locality:      "Paris",
	}

	// _, err := mapfitClient.GetDirections(srcAddr, dstAddr, "driving")
	_, err := mapfitClient.GetDirections(srcAddr, dstAddr, "walking")
	if err != nil {
		t.FailNow()
	}
}
