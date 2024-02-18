package swapi

import (
	"context"
	"encoding/json"
	"net/http"
)

type Client struct {
	provider *provider
}

func NewClient() *Client {
	return &Client{

		provider: newProvider(),
	}
}

func (c *Client) GetData(ctx context.Context, routeURL string, data interface{}) (err error) {
	request, err := c.provider.buildRequest(ctx, http.MethodGet, routeURL, nil)
	if err != nil {
		return err
	}

	var onSuccessFn SuccessFn = func(bodyByte []byte) interface{} {
		if err := json.Unmarshal(bodyByte, data); err != nil {
			return err
		}
		return nil
	}

	var onFailureFn FailureFn = func(response *http.Response, bodyByte []byte) error {
		errorResponse := ErrorResponse{}
		if err := json.Unmarshal(bodyByte, &errorResponse); err != nil {
			return err
		}
		return &errorResponse
	}
	result := HandleResponse(request, onSuccessFn, onFailureFn)
	if result.StatusCode == http.StatusNotFound {
		return nil
	}

	return result.Error
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Date    string `json:"datetime"`
}

func (e ErrorResponse) Error() string {
	return e.Message
}
