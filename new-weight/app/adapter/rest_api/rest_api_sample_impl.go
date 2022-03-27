package rest_api

import (
	"project-name/app/shared/utils"
)

func (r *restAPI) HitExternalAPI() (out map[string]interface{}, statusCode int, code string, err error) {
	payload := map[string]interface{}{
		"sample": 1,
	}

	params := utils.NewParams("http://localhost:8083/v1/un-decrypt", "", "", payload, nil, nil, "")
	out, statusCode, code, err = r.httpRequest.POSTWithoutBasicAuth(params)

	return
}