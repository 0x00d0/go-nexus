package v3

import (
	"net/http"
	"net/url"
)

type Nexus struct {
	url         string
	username    string
	password    string
	authHeader  string
	accessToken string
}

func NewNexus(url string, userName string, password string) *Nexus {
	return &Nexus{url: url, username: userName, password: password}
}

func (n *Nexus) do(endpoint string, responseStruct interface{}, options map[string]string) (*http.Response, error) {
	URL, err := url.Parse(n.url + endpoint + "/")
	if err != nil {
		return nil, err
	}
	request := NewAPIRequest(URL, n.username, n.password, n.authHeader, n.accessToken)

	return request.Get(responseStruct, options)
}
