package utils

import (
	"io"
	"log"
	"net/http"

	"github.com/slainsama/msgr_server/globals"
)

// HttpGET is a wrapper of http.Get
// It returns a *http.Response
func HttpGET(baseURL string, params map[string]string) (code int, body []byte, err error) {
	url := buildURL(baseURL, params)

	if isDEBUG := globals.UnmarshaledConfig.DEBUG.Switch; isDEBUG {
		log.Println("HttpGET:", url)
	}
	resp, err := http.Get(url)
	if err != nil {
		return 0, nil, err
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			log.Println(err)
		}
	}(resp.Body)

	// Return response body as []byte
	body, err = io.ReadAll(resp.Body)
	return resp.StatusCode, body, err
}

// buildURL builds a url with params
// If params is nil, it returns baseURL
func buildURL(baseURL string, params map[string]string) string {
	if params == nil {
		return baseURL
	}
	url := baseURL + "?"
	for key, value := range params {
		url += key + "=" + value + "&"
	}
	url = url[:len(url)-1]
	return url
}