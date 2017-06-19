// Package geoelevation provides elevation data from geonames service
package geoelevation

import (
	"encoding/json"
	"net/url"

	"github.com/variar/go-geonames-client/geoclient"
)

//Elevation structure
type Elevation struct {
	Height int     // height
	Lon    float32 // longitude
	Lat    float32 // latitude
}

func getElevationData(requester geoclient.Requester, endpoint string, lon string, lat string) ([]byte, error) {
	request := url.Values{}
	request.Set("lat", lat)
	request.Set("lng", lon)

	return requester.MakeRequest(endpoint, request)
}

//GetAstergdemElevation provides elevation for given longitude and latitude using astergdem data
//http://www.geonames.org/export/web-services.html#astergdem
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

//GetGtopo30Elevation provides elevation for given longitude and latitude using gtopo30 data
//http://www.geonames.org/export/web-services.html#gtopo30
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

//GetSrtm3Elevation provides elevation for given longitude and latitude using strm3 data
//http://www.geonames.org/export/web-services.html#strm3
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
