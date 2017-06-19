package geosearch

const (
	searchStyleShort  = "short"
	searchStyleMedium = "medium"
)

/*ShortGeoname data
{
      "toponymName": "London",
      "fcl": "P",
      "name": "London",
      "countryCode": "GB",
      "lng": "-0.12574",
      "fcode": "PPLC",
      "geonameId": 2643743,
      "lat": "51.50853"
}*/
type ShortGeoname struct {
	Lon          string `json:"lng"`
	Lat          string `json:"lat"`
	GeonameID    int64  `json:"geonameId"`
	CountryCode  string `json:"countryCode"`
	Name         string `json:"name"`
	ToponymName  string `json:"toponymName"`
	FeatureCode  string `json:"fcode"`
	FeatureClass string `json:"fcl"`
}

type shortSearchResponse struct {
	Results  int            `json:"totalResultsCount"`
	Geonames []ShortGeoname `json:"geonames"`
}

/*MediumGeoname data
{
  "countryId": "2635167",
  "adminCode1": "ENG",
  "countryName": "United Kingdom",
  "fclName": "city, village,...",
  "countryCode": "GB",
  "lng": "-0.12574",
  "fcodeName": "capital of a political entity",
  "toponymName": "London",
  "fcl": "P",
  "name": "London",
  "fcode": "PPLC",
  "geonameId": 2643743,
  "lat": "51.50853",
  "adminName1": "England",
  "population": 7556900
}*/
type MediumGeoname struct {
	ShortGeoname

	CountryID   string `json:"countryId"`
	CountryName string `json:"countryName"`
	CountryCode string `json:"countryCode"`

	AdminCode1 string `json:"adminCode1"`
	AdminName1 string `json:"adminName1"`

	FeatureCodeName  string `json:"fcodeName"`
	FeatureClassName string `json:"fclName"`

	Population int `json:"population"`
}

type mediumSearchResponse struct {
	Results  int             `json:"totalResultsCount"`
	Geonames []MediumGeoname `json:"geonames"`
}
