package geosearch

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/variar/go-geonames-client/geoclient"
)

/*NearbyQuery data to findNearby geonames API
http://www.geonames.org/export/web-services.html#findNearby
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

//NewNearbyQuery create new query with prefilled latitude and longitude parameters
func NewNearbyQuery(lon string, lat string) *NearbyQuery {
	return &NearbyQuery{Lat: lat, Lng: lon, MaxRows: 10}
}

//GetShortNearbyGeonames returns ShortGeoname for NearbyQuery
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

//GetMediumNearbyGeonames returns ShortGeoname for NearbyQuery
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
