package geosearch

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/variar/go-geonames-client/geoclient"
)

/*
Url : api.geonames.org/findNearby?
Parameters : lat,lng, featureClass,featureCode, radius: radius in km (optional), maxRows : max number of rows (default 10)
style : SHORT,MEDIUM,LONG,FULL (default = MEDIUM), verbosity of returned xml document
localCountry: in border areas this parameter will restrict the search on the local country, value=true
Result : returns the closest toponym for the lat/lng query as xml document
*/
type NearbyQuery struct {
	Lat string
	Lng string

	FeatureClasses []string
	FeatureCodes   []string

	Radius  int
	MaxRows int

	LocalCountry bool
}

func NewNearbyQuery(lon string, lat string) *NearbyQuery {
	return &NearbyQuery{Lat: lat, Lng: lon, MaxRows: 10}
}

func GetShortNearbyGeonames(requester geoclient.Requester, query *NearbyQuery) ([]ShortGeoname, error) {
	data, err := getNearbyGeonames(requester, query, searchStyleShort)
	if err != nil {
		return nil, err
	}

	var response shortSearchResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	return response.Geonames, nil
}

func GetMediumNearbyGeonames(requester geoclient.Requester, query *NearbyQuery) ([]MediumGeoname, error) {
	data, err := getNearbyGeonames(requester, query, searchStyleMedium)
	if err != nil {
		return nil, err
	}

	var response mediumSearchResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	return response.Geonames, nil
}

func getNearbyGeonames(requester geoclient.Requester, query *NearbyQuery, style string) ([]byte, error) {
	request := url.Values{}
	request.Set("style", style)
	request.Set("maxRows", strconv.Itoa(query.MaxRows))
	request.Set("lng", query.Lng)
	request.Set("lat", query.Lat)

	if query.Radius > 0 {
		request.Set("radius", strconv.Itoa(query.Radius))
	}

	if query.LocalCountry {
		request.Set("localCountry", "true")
	}

	for _, fClass := range query.FeatureClasses {
		request.Add("featureClass", fClass)
	}

	for _, fCode := range query.FeatureCodes {
		request.Add("featureCode", fCode)
	}

	return requester.MakeRequest("findNearbyJSON", request)
}
