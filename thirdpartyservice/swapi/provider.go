package swapi

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type provider struct {
	httpClient *http.Client
}

func newProvider() *provider {
	httpClient := &http.Client{}
	return &provider{
		httpClient: httpClient,
	}
}

func (provider *provider) buildRequest(ctx context.Context, method, url string, rawBody interface{}) (*http.Request, error) {
	var request *http.Request
	var err error
	if rawBody != nil {
		jsonStr, err := json.Marshal(rawBody)
		if err != nil {
			return nil, err
		}
		log.Println("JSON Body : ", string(jsonStr))
		body := bytes.NewBuffer(jsonStr)
		request, err = http.NewRequest(method, url, body)
		request = request.WithContext(ctx)
		if err != nil {
			return nil, err
		}
		request.Header.Set("Content-Type", "application/json")
	} else {
		request, err = http.NewRequest(method, url, nil)
		if err != nil {
			return nil, err
		}
	}

	return request, nil
}
