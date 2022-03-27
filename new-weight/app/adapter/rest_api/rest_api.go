package rest_api

type RestAPI interface {
	HitExternalAPI() (out map[string]interface{}, statusCode int, code string, err error)
}