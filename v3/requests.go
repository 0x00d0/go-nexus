package v3

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type APIRequest struct {
	Url         *url.URL
	UserName    string
	Password    string
	Header      http.Header
	Client      *http.Client
	timeout     time.Duration
	Payload     io.Reader
	authHeader  string
	accessToken string
}

func NewAPIRequest(url *url.URL, userName, password, authHeader, accessToken string) *APIRequest {
	return &APIRequest{Url: url, UserName: userName, Password: password, authHeader: authHeader, accessToken: accessToken, Client: &http.Client{}}
}

func (r *APIRequest) SetHeader(key string, value string) *APIRequest {
	r.Header.Set(key, value)
	return r
}

func (r *APIRequest) Get(responseStruct interface{}, query map[string]string) (*http.Response, error) {
	var header = http.Header{}
	r.Header = header
	r.SetHeader("Content-Type", "application/json")
	return r.Do("GET", responseStruct, query)
}

func (r *APIRequest) Post(payload io.Reader, responseStruct interface{}, query map[string]string) (*http.Response, error) {
	var header = http.Header{}
	r.Header = header
	r.Payload = payload
	r.SetHeader("Content-Type", "application/json")
	return r.Do("POST", responseStruct, query)
}

func (r *APIRequest) Do(method string, responseStruct interface{}, options ...interface{}) (*http.Response, error) {

	for _, option := range options {
		switch v := option.(type) {
		case map[string]string:
			query := make(url.Values)
			for key, val := range v {
				query.Set(key, val)
			}
			r.Url.RawQuery = query.Encode()
		}
	}
	request, err := http.NewRequest(method, r.Url.String(), r.Payload)
	if err != nil {
		return nil, err
	}

	for k := range r.Header {
		r.Header.Add(k, r.Header.Get(k))
	}

	if r.accessToken != "" {
		if r.authHeader != "" {
			request.Header.Set(r.authHeader, r.accessToken)
		} else {
			request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", r.accessToken))
		}
	} else {
		request.SetBasicAuth(r.UserName, r.Password)
	}
	response, err := r.Client.Do(request)
	if err != nil {
		return nil, err
	}

	//if response.StatusCode != 200 {
	//	return nil, errors.New(response.Status)
	//}

	return r.JSONResponse(response, responseStruct)
}

func (r *APIRequest) JSONResponse(response *http.Response, responseStruct interface{}) (*http.Response, error) {
	defer response.Body.Close()
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(b, responseStruct)
	if err != nil {
		return response, err
	}
	return response, nil
}
