/*Package geotimezone provides timezone data from coordinates
http://www.geonames.org/export/web-services.html#timezone
{
  "sunrise": "2016-06-09 05:24",
  "lng": 10.2,
  "countryCode": "AT",
  "gmtOffset": 1,
  "rawOffset": 1,
  "sunset": "2016-06-09 21:13",
  "timezoneId": "Europe/Vienna",
  "dstOffset": 2,
  "countryName": "Austria",
  "time": "2016-06-09 00:36",
  "lat": 47.01
}
*/
package geotimezone

import (
	"encoding/json"
	"net/url"

	"github.com/variar/go-geonames-client/geoclient"
)

// Timezone data as reported by geonames timezone JSON service
// http://www.geonames.org/export/web-services.html#timezone
type Timezone struct {
	TimezoneId  string `json:"timezoneId"`
	CountryCode string `json:"countryCode"`
	CountryName string `json:"countryName"`

	GmtOffset int `json:"gmtOffset"`
	RawOffset int `json:"rawOffset"`
	DstOffset int `json:"dstOffset"`

	Lat float32 `json:"lat"`
	Lon float32 `json:"lng"`
}

// GetTimezone returns timezone data for given longitude and latitude
func GetTimezone(requester geoclient.Requester, lon string, lat string) (Timezone, error) {
	request := url.Values{}
	request.Set("lat", lat)
	request.Set("lng", lon)

	data, err := requester.MakeRequest("timezoneJSON", request)
	if err != nil {
		return Timezone{}, err
	}

	var tz Timezone
	err = json.Unmarshal(data, &tz)
	if err != nil {
		return Timezone{}, err
	}

	return tz, nil
}
