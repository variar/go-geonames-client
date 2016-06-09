package geoelevation

import (
	"encoding/json"
	"net/url"

	"github.com/variar/go-geonames-client/geoclient"
)

type Elevation struct {
	Height int
	Lon    float32
	Lat    float32
}

func getElevationData(requester geoclient.Requester, endpoint string, lon string, lat string) ([]byte, error) {
	request := url.Values{}
	request.Set("lat", lat)
	request.Set("lng", lon)

	return requester.MakeRequest(endpoint, request)
}

func GetAstergdemElevation(requester geoclient.Requester, lon string, lat string) (Elevation, error) {
	data, err := getElevationData(requester, "astergdemJSON", lon, lat)
	if err != nil {
		return Elevation{}, err
	}

	type astergdem struct {
		Lng       float32
		Astergdem int
		Lat       float32
	}
	var result astergdem
	err = json.Unmarshal(data, &result)
	if err != nil {
		return Elevation{}, err
	}

	return Elevation{Height: result.Astergdem, Lat: result.Lat, Lon: result.Lng}, nil
}

func GetGtopo30Elevation(requester geoclient.Requester, lon string, lat string) (Elevation, error) {
	data, err := getElevationData(requester, "gtopo30JSON", lon, lat)
	if err != nil {
		return Elevation{}, err
	}

	type gtopo30 struct {
		Lng     float32
		Gtopo30 int
		Lat     float32
	}
	var result gtopo30
	err = json.Unmarshal(data, &result)
	if err != nil {
		return Elevation{}, err
	}

	return Elevation{Height: result.Gtopo30, Lat: result.Lat, Lon: result.Lng}, nil
}

func GetSrtm3Elevation(requester geoclient.Requester, lon string, lat string) (Elevation, error) {
	data, err := getElevationData(requester, "srtm3JSON", lon, lat)
	if err != nil {
		return Elevation{}, err
	}

	type srtm3 struct {
		Lng   float32
		Srtm3 int
		Lat   float32
	}
	var result srtm3
	err = json.Unmarshal(data, &result)
	if err != nil {
		return Elevation{}, err
	}

	return Elevation{Height: result.Srtm3, Lat: result.Lat, Lon: result.Lng}, nil
}
