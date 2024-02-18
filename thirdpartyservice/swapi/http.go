package swapi

import (
	"io"
	"net/http"
)

// Result struct
type Result struct {
	Error      error
	Data       interface{}
	StatusCode int
}

// SuccessFn type
type SuccessFn func(bodyByte []byte) interface{}

// FailureFn type
type FailureFn func(response *http.Response, bodyByte []byte) error

// HandleResponse func
func HandleResponse(request *http.Request, successFn SuccessFn, failureFn FailureFn) *Result {
	result := &Result{}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		result.Error = err

		return result
	}
	defer response.Body.Close()

	bodyByte, err := io.ReadAll(response.Body)
	if err != nil {
		result.Error = err
		return result
	}
	result.StatusCode = response.StatusCode
	if response.StatusCode >= http.StatusOK && response.StatusCode <= http.StatusMultipleChoices {
		result.Data = successFn(bodyByte)
	} else {
		result.Error = failureFn(response, bodyByte)
	}

	return result
}
