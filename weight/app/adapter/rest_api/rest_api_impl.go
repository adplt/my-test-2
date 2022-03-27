package rest_api

import "project-name/app/shared/utils"

type restAPI struct {
	http *utils.HttpUtils
	httpRequest *utils.UtilHttpRequest
}

func NewRestAPI(http *utils.HttpUtils, httpRequest *utils.UtilHttpRequest) RestAPI {
	return &restAPI{http: http, httpRequest: httpRequest}
}