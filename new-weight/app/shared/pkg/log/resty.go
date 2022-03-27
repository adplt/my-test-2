package log

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

// OnBeforeRequest godoc.
func OnBeforeRequest(ctx *gin.Context) func(client *resty.Client, request *resty.Request) error {
	return func(client *resty.Client, request *resty.Request) error {
		if ctx == nil {
			return errors.New("context is null")
		}

		msg := "Request  "
		if request != nil {
			msg += fmt.Sprintf("%v %v ", request.Method, request.URL)
			msg += parseToString(request.Body)
		}

		Info(ctx, msg)
		return nil
	}
}

// OnAfterResponse godoc.
func OnAfterResponse(ctx *gin.Context) func(client *resty.Client, response *resty.Response) error {
	return func(client *resty.Client, response *resty.Response) error {
		if ctx == nil {
			return errors.New("context is null")
		}

		msg := "Response "
		if response != nil {
			msg += fmt.Sprintf("%v %v %v %v ", response.Request.Method, response.Request.URL, response.StatusCode(), response.Time())
			msg += parseToString(response.Body())
		}

		Info(ctx, msg)
		return nil
	}
}

func parseToString(in interface{}) string {
	if in == nil {
		return ""
	}

	var body map[string]interface{}

	switch in.(type) {
	case []byte:
		_ = json.Unmarshal(in.([]byte), &body)
	default:
		b, err := json.MarshalIndent(in, "", " ")
		if err == nil {
			_ = json.Unmarshal(b, &body)
		}
	}

	return fmt.Sprintf("%+v", body)
}
