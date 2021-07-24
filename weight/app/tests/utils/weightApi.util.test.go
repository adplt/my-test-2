package testutil

import (
	"net/http"

	httputils "project/app/utils/http"
	variables "project/app/variables"
)

// GetWeightAPI function
func GetWeightAPI(query string) (*http.Response, *httputils.CommonJSON) {
	server := ServerTest()
	defer server.Close()
	return httputils.GET(server.URL+"/"+variables.API_VERSION+"/"+variables.WEIGHT, query)
}

// AddWeightAPI function
func AddWeightAPI(payload map[string]interface{}) (*http.Response, *httputils.CommonJSON) {
	server := ServerTest()
	defer server.Close()
	return httputils.POST(server.URL+"/"+variables.API_VERSION+"/"+variables.WEIGHT, "", payload)
}

// UpdateWeightAPI function
func UpdateWeightAPI(id string, payload map[string]interface{}) (*http.Response, *httputils.CommonJSON) {
	server := ServerTest()
	defer server.Close()
	return httputils.PUT(server.URL+"/"+variables.API_VERSION+"/"+variables.WEIGHT, "/"+id, payload)
}
