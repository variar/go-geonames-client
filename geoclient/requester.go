// Package geoclient provides an interface to make requests to geonames API
package geoclient

import (
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/golang/glog"
)

// Requester interface is used to make a request to provided endpoint
type Requester interface {
	MakeRequest(endpoint string, params url.Values) ([]byte, error)
}

type basicRequester struct {
	username string
	lang     string
}

const (
	geonamesURL = "http://api.geonames.org/"
)

func (requester *basicRequester) MakeRequest(endpoint string, params url.Values) ([]byte, error) {
	params.Set("username", requester.username)
	params.Set("lang", requester.lang)

	query := geonamesURL + endpoint + "?" + params.Encode()
	glog.Infoln("Query:", query)

	resp, err := http.Get(query)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	glog.V(1).Infof("Response: %s\n", data)

	return data, nil
}

// NewRequester creates a requester with preset username and language parameters
func NewRequester(username string, lang string) Requester {
	return &basicRequester{username: username, lang: lang}
}
