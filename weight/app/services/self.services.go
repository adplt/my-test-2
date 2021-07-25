package services

import (
	"net/http"
	configs "project/app/configs"
	httputils "project/app/utils/http"
	variables "project/app/variables"
)

// GetWeightAPI function
func GetWeightAPI(date string) (*http.Response, *httputils.CommonJSON) {
	if date != "" {
		date = "?date=" + date
	}
	httpRes, err := httputils.GET(configs.Env.AppURL+"/"+variables.API_VERSION+"/"+variables.WEIGHT, date+"/")
	return httpRes, err
}

// GetWeightByIdAPI function
func GetWeightByIdAPI(id string) (*http.Response, *httputils.CommonJSON) {
	httpRes, err := httputils.GET(configs.Env.AppURL+"/"+variables.API_VERSION+"/"+variables.WEIGHT, "/"+id+"/")
	return httpRes, err
}

// AddWeightAPI function
func AddWeightAPI(payload map[string]interface{}) (*http.Response, *httputils.CommonJSON) {
	httpRes, err := httputils.POST(configs.Env.AppURL+"/"+variables.API_VERSION+"/"+variables.WEIGHT+"/", "", payload)
	return httpRes, err
}

// UpdateWeightAPI function
func UpdateWeightAPI(id string, payload map[string]interface{}) (*http.Response, *httputils.CommonJSON) {
	httpRes, err := httputils.PUT(configs.Env.AppURL+"/"+variables.API_VERSION+"/"+variables.WEIGHT, "/"+id+"/", payload)
	return httpRes, err
}
