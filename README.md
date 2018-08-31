[![GoDoc](https://godoc.org/github.com/albttx/go-mapfit?status.svg)](https://godoc.org/github.com/albttx/go-mapfit)
[![CircleCI](https://circleci.com/gh/albttx/go-mapfit/tree/master.svg?style=svg)](https://circleci.com/gh/albttx/go-mapfit/tree/master)

# Mapfit Golang client (non-official)

## Mapfit REST APIS

- [x] Directions [docs](https://docs.mapfit.com/docs/rest_directions-api)
- [x] Geocoding [docs](https://docs.mapfit.com/docs/rest_geocoding-api)

## Exemple

```go
mapfitClient := mapfit.NewClient("mapfit-secret-token")
srcAddr := mapfit.Address{
    StreetAddress: "avenue des Champs Elysee",
    Locality:      "Paris",
}
dstAddr := mapfit.Address{
    StreetAddress: "Tour Eiffel",
    Locality:      "Paris",
}

// Directions API
directions, err := mapfitClient.GetDirections(srcAddr, dstAddr, "driving")
if err != nil {
    panic(err)
}
fmt.Printf("%#v\n", directions)

// Geocode API
geocode, err := mapfitClient.Geocode(srcAddr)
if err != nil {
    panic(err)
}
fmt.Printf("%#v\n", geocode)
```