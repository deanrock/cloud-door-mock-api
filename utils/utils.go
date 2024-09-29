package utils

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strings"

	"github.com/deanrock/cloud-door-mock-api/client"
)

func Pointer[T any](x T) *T {
	return &x
}

func IsSlice(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Slice
}

func IsFormEncoded(req *http.Request) bool {
	contentType := req.Header.Get("Content-Type")
	return strings.Contains(contentType, "application/x-www-form-urlencoded")
}

// Result can be either an object or array. OpenAPI type is wrong.
type AbpWebModelsAjaxResponse struct {
	Abp                 *bool                         `json:"__abp,omitempty"`
	Error               *client.AbpWebModelsErrorInfo `json:"error,omitempty"`
	Result              *interface{}                  `json:"result,omitempty"`
	Success             *bool                         `json:"success,omitempty"`
	TargetUrl           *string                       `json:"targetUrl,omitempty"`
	UnAuthorizedRequest *bool                         `json:"unAuthorizedRequest,omitempty"`
}

func ToAbpResponse(data interface{}) AbpWebModelsAjaxResponse {
	encoded, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	var decoded interface{}
	err = json.Unmarshal(encoded, &decoded)
	if err != nil {
		panic(err)
	}

	return AbpWebModelsAjaxResponse{
		Abp:                 Pointer(true),
		Error:               nil,
		Result:              Pointer(decoded),
		Success:             Pointer(true),
		TargetUrl:           nil,
		UnAuthorizedRequest: Pointer(false),
	}
}
