package geosearch

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/variar/go-geonames-client/geoclient"
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

type SearchQuery struct {
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

func NewSearchQuery(searchType string, name string) *SearchQuery {
	return &SearchQuery{Fuzzy: 1, SearchType: searchType, Query: name, MaxRows: 10}
}

func GetShortGeonames(requester geoclient.Requester, query *SearchQuery) ([]ShortGeoname, error) {
	data, err := getGeonames(requester, query, searchStyleShort)
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

func GetMediumGeonames(requester geoclient.Requester, query *SearchQuery) ([]MediumGeoname, error) {
	data, err := getGeonames(requester, query, searchStyleMedium)
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

func getGeonames(requester geoclient.Requester, query *SearchQuery, style string) ([]byte, error) {
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
