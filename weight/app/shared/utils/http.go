package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"project-name/app/shared/config"
	"project-name/app/shared/pkg/log"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/gomodul/fn"
)

type HttpUtils struct {
	client  *resty.Client
	request *resty.Request
}

func NewUtilsHttp() *HttpUtils {
	client := NewClient()
	return &HttpUtils{
		client: client,
	}
}

func (q *HttpUtils) GetClient() *http.Client {
	return q.client.GetClient()
}

func (q *HttpUtils) NewRequest(ctx *gin.Context) *resty.Request {
	var flag = "OnBeforeRequestKey"
	if !ctx.GetBool(flag) {
		q.client.OnBeforeRequest(log.OnBeforeRequest(ctx))
		ctx.Set(flag, true)
	}

	flag = "OnAfterResponseKey"
	if !ctx.GetBool(flag) {
		q.client.OnAfterResponse(log.OnAfterResponse(ctx))
		ctx.Set(flag, true)
	}

	q.request = q.client.R()
	return q.request
}

// Do godoc.
func (q *HttpUtils) Do(ctx *gin.Context, result interface{}) (statusCode int, code string, err error) {
	var res *resty.Response
	if res, err = q.request.Send(); err != nil {
		log.Error(ctx, err)
		if e, ok := err.(net.Error); ok && e.Timeout() {
			statusCode = http.StatusGatewayTimeout
			code = config.RCMobeGatewayTimeout
		} else {
			statusCode = http.StatusInternalServerError
			code = config.RCMobeInternalServerError
		}
		return
	}

	if res.StatusCode() < http.StatusOK || res.StatusCode() >= http.StatusMultipleChoices {
		err = errors.New("http.ResponseCode: " + http.StatusText(res.StatusCode()))
		log.Error(ctx, err)

		statusCode = res.StatusCode()
		code = config.RCMobeInternalServerError
		return
	}

	if res.RawBody() != nil {
		defer func() { fn.Check(res.RawBody().Close) }()
		if err = json.Unmarshal(res.Body(), &result); err != nil {
			err = errors.New("json.Unmarshal(res.Body(), &result): " + err.Error())
			if e, ok := err.(*json.SyntaxError); ok {
				err = fmt.Errorf("%v at byte %v", err.Error(), e.Offset)
			}
			log.Error(ctx, err)

			statusCode = http.StatusInternalServerError
			code = config.RCMobeInternalServerError
			return
		}
	}

	return http.StatusOK, config.RCMobeOK, nil
}

func (q *HttpUtils) Request(ctx *gin.Context, method, url string, body, result interface{}) (int, string, error) {
	q.NewRequest(ctx).SetBody(body)
	q.request.Method = method
	q.request.URL = url
	return q.Do(ctx, result)
}

func NewClient() *resty.Client {
	client := resty.New()
	client.SetHeader("User-Agent", "MNC_Bank_MB_Engine/1.0")
	client.SetTimeout(config.Request.Timeout)
	return client
}
