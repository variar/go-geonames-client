package geosearch

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/variar/go-geonames-client/geoclient"
)

/*
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

/*{
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
},*/

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

const (
	SearchStyleShort  = "short"
	SearchStyleMedium = "medium"
)

const (
	QueryName           = "name"
	QueryEquals         = "name_equals"
	QueryNameStartsWith = "name_startsWith"
	QueryAll            = "q"
)

const (
	Cities1000  = "cities1000"
	Cities5000  = "cities5000"
	Cities15000 = "cities15000"
)

const (
	OrderByRelevance  = "relevance"
	OrderByPopulation = "population"
	OrderByElevation  = "elevation"
)

type GeonamesQuery struct {
	Query      string
	SearchType string

	IsNameRequired bool

	SearchLang string

	MaxRows  int
	StartRow int

	CountryCodes  []string
	CountryBias   string
	ContinentCode string

	AdminCode1 string
	AdminCode2 string
	AdminCode3 string

	FeatureClasses []string
	FeatureCodes   []string

	CitiesSize string

	Tag   string
	Fuzzy float64

	LimitBoundingBox         bool
	East, West, North, South float64

	OrderBy string

	IncludeBoundingBox bool
}

func NewQuery(searchType string, name string) GeonamesQuery {
	return GeonamesQuery{Fuzzy: 1, SearchType: searchType, Query: name, MaxRows: 10}
}

func GetShortGeonames(requester geoclient.Requester, query GeonamesQuery) ([]ShortGeoname, error) {
	data, err := getGeonames(requester, query, SearchStyleShort)
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

func GetMediumGeonames(requester geoclient.Requester, query GeonamesQuery) ([]MediumGeoname, error) {
	data, err := getGeonames(requester, query, SearchStyleMedium)
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

func getGeonames(requester geoclient.Requester, query GeonamesQuery, style string) ([]byte, error) {
	request := url.Values{}
	request.Set("style", style)
	request.Set(query.SearchType, query.Query)
	request.Set("maxRows", strconv.Itoa(query.MaxRows))
	request.Set("startRow", strconv.Itoa(query.StartRow))

	for _, country := range query.CountryCodes {
		request.Add("country", country)
	}

	if len(query.CountryBias) > 0 {
		request.Set("countryBias", query.CountryBias)
	}

	if len(query.ContinentCode) > 0 {
		request.Set("continentCode", query.ContinentCode)
	}

	if len(query.AdminCode1) > 0 {
		request.Set("adminCode1", query.AdminCode1)
	}
	if len(query.AdminCode2) > 0 {
		request.Set("adminCode2", query.AdminCode2)
	}
	if len(query.AdminCode3) > 0 {
		request.Set("adminCode3", query.AdminCode3)
	}

	for _, fClass := range query.FeatureClasses {
		request.Add("featureClass", fClass)
	}

	for _, fCode := range query.FeatureCodes {
		request.Add("featureCode", fCode)
	}

	if len(query.CitiesSize) > 0 {
		request.Set("cities", query.CitiesSize)
	}

	if query.IsNameRequired {
		request.Set("isNameRequired", "true")
	}
	if query.IncludeBoundingBox {
		request.Set("inclBbox", "true")
	}

	if query.SearchType == QueryName {
		request.Set("fuzzy", strconv.FormatFloat(query.Fuzzy, 'f', -1, 32))
	}

	if query.LimitBoundingBox {
		request.Set("east", strconv.FormatFloat(query.East, 'f', -1, 32))
		request.Set("west", strconv.FormatFloat(query.West, 'f', -1, 32))
		request.Set("north", strconv.FormatFloat(query.North, 'f', -1, 32))
		request.Set("south", strconv.FormatFloat(query.South, 'f', -1, 32))
	}

	if len(query.OrderBy) > 0 {
		request.Set("orderBy", query.OrderBy)
	}

	if len(query.SearchLang) > 0 {
		request.Set("searchlang", query.SearchLang)
	}

	return requester.MakeRequest("searchJSON", request)
}
