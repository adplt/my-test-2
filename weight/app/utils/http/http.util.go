package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	configs "project/app/configs"
	"time"
)

var client = &http.Client{Timeout: time.Second * time.Duration(30)}

// CheckResponse function
func CheckResponse(res *http.Response) (*http.Response, *CommonJSON) {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		configs.Throw(err)
	}

	result := map[string]interface{}{}
	if err := json.Unmarshal(body, &result); err != nil {
		configs.Throw(err)
	}

	switch res.StatusCode {
	case http.StatusBadRequest:
		return nil, Error400(errors.New("Bad Request"), body)
	case http.StatusInternalServerError:
		return nil, Error500(errors.New("Internal Server Error"), body)
	}

	return res, nil
}

// GET function
func GET(URL string, params string) (*http.Response, *CommonJSON) {
	req, err := http.NewRequest(http.MethodGet, URL+params, nil)
	if err != nil {
		configs.Throw(err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		configs.Throw(err)
	}

	if res != nil && res.StatusCode != http.StatusOK {
		return CheckResponse(res)
	}

	return res, nil
}

// POST function
func POST(URL string, params string, reqPayload map[string]interface{}) (*http.Response, *CommonJSON) {
	payload, err := json.Marshal(reqPayload)
	if err != nil {
		configs.Throw(err)
	}
	req, err := http.NewRequest(http.MethodPost, URL+params, bytes.NewBuffer(payload))

	if err != nil {
		configs.Throw(err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		configs.Throw(err)
	}

	if res != nil && res.StatusCode != http.StatusOK {
		return CheckResponse(res)
	}

	return res, nil
}

// PUT function
func PUT(URL string, params string, reqPayload map[string]interface{}) (*http.Response, *CommonJSON) {
	payload, err := json.Marshal(reqPayload)
	if err != nil {
		configs.Throw(err)
	}
	req, err := http.NewRequest(http.MethodPut, URL+params, bytes.NewBuffer(payload))
	if err != nil {
		configs.Throw(err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		configs.Throw(err)
	}

	if res != nil && res.StatusCode != http.StatusOK {
		return CheckResponse(res)
	}

	return res, nil
}

// DELETE function
func DELETE(URL string, params string, reqPayload map[string]interface{}) (*http.Response, *CommonJSON) {
	payload, err := json.Marshal(reqPayload)
	if err != nil {
		configs.Throw(err)
	}
	req, err := http.NewRequest(http.MethodDelete, URL+params, bytes.NewBuffer(payload))
	if err != nil {
		configs.Throw(err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		configs.Throw(err)
	}

	if res != nil && res.StatusCode != http.StatusOK {
		return CheckResponse(res)
	}

	return res, nil
}

// POST_ARRAY function
func POST_ARRAY(URL string, params string, reqPayload []map[string]interface{}) (*http.Response, *CommonJSON) {
	payload, err := json.Marshal(reqPayload)
	if err != nil {
		configs.Throw(err)
	}
	req, err := http.NewRequest(http.MethodPost, URL+params, bytes.NewBuffer(payload))

	if err != nil {
		configs.Throw(err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		configs.Throw(err)
	}

	if res != nil && res.StatusCode != http.StatusOK {
		return CheckResponse(res)
	}

	return res, nil
}

// PUT_ARRAY function
func PUT_ARRAY(URL string, params string, reqPayload []map[string]interface{}) (*http.Response, *CommonJSON) {
	payload, err := json.Marshal(reqPayload)
	if err != nil {
		configs.Throw(err)
	}
	req, err := http.NewRequest(http.MethodPut, URL+params, bytes.NewBuffer(payload))
	if err != nil {
		configs.Throw(err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		configs.Throw(err)
	}

	if res != nil && res.StatusCode != http.StatusOK {
		return CheckResponse(res)
	}

	return res, nil
}

// DELETE_ARRAY function
func DELETE_ARRAY(URL string, params string, reqPayload []map[string]interface{}) (*http.Response, *CommonJSON) {
	payload, err := json.Marshal(reqPayload)
	if err != nil {
		configs.Throw(err)
	}
	req, err := http.NewRequest(http.MethodDelete, URL+params, bytes.NewBuffer(payload))
	if err != nil {
		configs.Throw(err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		configs.Throw(err)
	}

	if res != nil && res.StatusCode != http.StatusOK {
		return CheckResponse(res)
	}

	return res, nil
}
