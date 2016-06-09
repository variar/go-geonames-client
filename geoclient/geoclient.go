package geoclient

import (
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/golang/glog"
)

type Requester interface {
	MakeRequest(endpoint string, params url.Values) ([]byte, error)
}

type Client struct {
	username string
	lang     string
}

const (
	geonamesURL = "http://api.geonames.org/"
)

func (client *Client) MakeRequest(endpoint string, params url.Values) ([]byte, error) {
	params.Set("username", client.username)
	params.Set("lang", client.lang)

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

func NewClient(username string, lang string) Requester {
	return &Client{username: username, lang: lang}
}
