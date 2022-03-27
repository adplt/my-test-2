package utils

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"project-name/app/shared/config"
	"time"

	"github.com/go-resty/resty/v2"
)

type UtilHttpRequest struct {
	GET                  func(params Params) (out map[string]interface{}, statusCode int, code string, err error)
	POST                 func(params Params) (out map[string]interface{}, statusCode int, code string, err error)
	PUT                  func(params Params) (out map[string]interface{}, statusCode int, code string, err error)
	POSTWithoutBasicAuth func(params Params) (out map[string]interface{}, statusCode int, code string, err error)
}

func NewUtilHttpRequest() *UtilHttpRequest {
	return &UtilHttpRequest{
		GET:                  GET,
		POST:                 POST,
		PUT:                  PUT,
		POSTWithoutBasicAuth: POSTWithoutBasicAuth,
	}
}

type Params struct {
	Url               string
	BasicAuthUsername string
	BasicAuthPassword string
	RequestBody       map[string]interface{}
	RequestParams     map[string]string
	RequestHeader     map[string]string
	Resty             *resty.Client
	BearerToken       string
}

func NewParams(url string, basicAuthUsername string, basicAuthPassword string, requestBody map[string]interface{},
	requestParams map[string]string, requestHeader map[string]string, bearerToken string) (out Params) {
	out.Url = url
	out.BasicAuthUsername = basicAuthUsername
	out.BasicAuthPassword = basicAuthPassword
	out.RequestBody = requestBody
	out.RequestParams = requestParams
	out.RequestHeader = requestHeader
	out.Resty = resty.New()
	out.BearerToken = bearerToken
	return out
}

func GET(params Params) (out map[string]interface{}, statusCode int, code string, err error) {
	params.Resty.SetTimeout(time.Second * 10)
	params.Resty.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	params.Resty.SetDebug(true)
	host := fmt.Sprint(params.Url)
	resp, err := params.Resty.R().
		SetBasicAuth(params.BasicAuthUsername, params.BasicAuthPassword).
		SetQueryParams(params.RequestParams).
		SetHeaders(params.RequestHeader).
		Get(host)
	if err != nil {
		return
	}

	_ = json.Unmarshal(resp.Body(), &out)

	if e, ok := err.(net.Error); ok && e.Timeout() {
		statusCode = http.StatusGatewayTimeout
		code = config.RCMobeGatewayTimeout
	} else {
		statusCode = http.StatusInternalServerError
		code = config.RCMobeInternalServerError
	}

	if resp.StatusCode() < http.StatusOK || resp.StatusCode() >= http.StatusMultipleChoices {
		err = errors.New("http.ResponseCode: " + http.StatusText(resp.StatusCode()))

		statusCode = resp.StatusCode()
		code = config.RCMobeInternalServerError
		return
	}
	return
}

func POST(params Params) (out map[string]interface{}, statusCode int, code string, err error) {
	params.Resty.SetTimeout(time.Second * 10)
	params.Resty.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	params.Resty.SetDebug(true)
	payload, _ := json.Marshal(params.RequestBody)
	host := fmt.Sprint(params.Url)
	resp, err := params.Resty.R().
		SetBasicAuth(params.BasicAuthUsername, params.BasicAuthPassword).
		SetQueryParams(params.RequestParams).
		SetBody(bytes.NewBuffer(payload)).
		Post(host)
	if err != nil {
		return
	}

	_ = json.Unmarshal(resp.Body(), &out)

	if e, ok := err.(net.Error); ok && e.Timeout() {
		statusCode = http.StatusGatewayTimeout
		code = config.RCMobeGatewayTimeout
	} else {
		statusCode = http.StatusInternalServerError
		code = config.RCMobeInternalServerError
	}

	if resp.StatusCode() < http.StatusOK || resp.StatusCode() >= http.StatusMultipleChoices {
		err = errors.New("http.ResponseCode: " + http.StatusText(resp.StatusCode()))

		statusCode = resp.StatusCode()
		code = config.RCMobeInternalServerError
		return
	}

	return
}

func PUT(params Params) (out map[string]interface{}, statusCode int, code string, err error) {
	params.Resty.SetTimeout(time.Second * 10)
	params.Resty.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	params.Resty.SetDebug(true)
	payload, _ := json.Marshal(params.RequestBody)
	host := fmt.Sprint(params.Url)
	var resp *resty.Response
	resp, err = params.Resty.R().
		SetBasicAuth(params.BasicAuthUsername, params.BasicAuthPassword).
		SetBody(bytes.NewBuffer(payload)).
		SetHeaders(params.RequestHeader).
		Put(host)
	if err != nil {
		return
	}

	_ = json.Unmarshal(resp.Body(), &out)

	if e, ok := err.(net.Error); ok && e.Timeout() {
		statusCode = http.StatusGatewayTimeout
		code = config.RCMobeGatewayTimeout
	} else {
		statusCode = http.StatusInternalServerError
		code = config.RCMobeInternalServerError
	}

	if resp.StatusCode() < http.StatusOK || resp.StatusCode() >= http.StatusMultipleChoices {
		err = errors.New("http.ResponseCode: " + http.StatusText(resp.StatusCode()))

		statusCode = resp.StatusCode()
		code = config.RCMobeInternalServerError
		return
	}

	return
}

func POSTWithoutBasicAuth(params Params) (out map[string]interface{}, statusCode int, code string, err error) {
	params.Resty.SetTimeout(time.Second * 10)
	params.Resty.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	params.Resty.SetDebug(true)
	payload, _ := json.Marshal(params.RequestBody)
	host := fmt.Sprint(params.Url)
	var resp *resty.Response

	resp, err = params.Resty.R().
		SetBody(bytes.NewBuffer(payload)).
		Post(host)
	if err != nil {
		return
	}

	_ = json.Unmarshal(resp.Body(), &out)

	if e, ok := err.(net.Error); ok && e.Timeout() {
		statusCode = http.StatusGatewayTimeout
		code = config.RCMobeGatewayTimeout
	} else {
		statusCode = http.StatusInternalServerError
		code = config.RCMobeInternalServerError
	}

	if resp.StatusCode() < http.StatusOK || resp.StatusCode() >= http.StatusMultipleChoices {
		err = errors.New("http.ResponseCode: " + http.StatusText(resp.StatusCode()))

		statusCode = resp.StatusCode()
		code = config.RCMobeInternalServerError
		return
	}
	return
}
