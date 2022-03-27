package sample

import (
	structs "project-name/app/shared/structs"
	"project-name/app/usecase/sample/response"
)

type UseCase interface {
	GetWeightById(in structs.GetWeight) (out response.WeightResponse, httpCode int, code string, err error)
	GetWeight(in structs.GetWeight) (out []response.WeightResponse, httpCode int, code string, err error)
	AddWeight(in structs.WeightData) (out response.WeightResponse, httpCode int, code string, err error)
	UpdateWeight(in structs.WeightData) (out response.WeightResponse, httpCode int, code string, err error)
}
