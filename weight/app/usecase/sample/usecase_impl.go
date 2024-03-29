package sample

import (
	"encoding/json"
	restApi "project-name/app/adapter/rest_api"
	"project-name/app/models/sql"
	repoWeight "project-name/app/repository/weight"
	structs "project-name/app/shared/structs"
	"project-name/app/usecase/sample/response"
)

type sampleUseCase struct {
	restAPI  restApi.RestAPI
	weightUC repoWeight.RepositoryWeight
}

func NewSampleUseCase(restAPI restApi.RestAPI, weightUC repoWeight.RepositoryWeight) UseCase {
	return &sampleUseCase{
		restAPI:  restAPI,
		weightUC: weightUC,
	}
}

func (uc *sampleUseCase) GetWeightById(in structs.GetWeight) (out response.WeightResponse, httpCode int, code string, err error) {
	res, err := uc.weightUC.FindById(*in.WeightRecordId)

	// convert map to json
	jsonString, _ := json.Marshal(res)
	json.Unmarshal(jsonString, &out)
	return
}

func (uc *sampleUseCase) GetWeight(in structs.GetWeight) (out []response.WeightResponse, httpCode int, code string, err error) {
	res, err := uc.weightUC.FindAll(in.Limit, in.Offset, "")

	// convert map to json
	jsonString, _ := json.Marshal(res)
	json.Unmarshal(jsonString, &out)
	return
}

func (uc *sampleUseCase) AddWeight(in structs.WeightData) (out response.WeightResponse, httpCode int, code string, err error) {
	var model sql.TxWeightRecord

	inJSON, _ := json.Marshal(in)
	json.Unmarshal(inJSON, &model)

	res, err := uc.weightUC.SaveWeight(model)

	// convert map to json
	jsonString, _ := json.Marshal(res)
	json.Unmarshal(jsonString, &out)
	return
}

func (uc *sampleUseCase) UpdateWeight(in structs.WeightData) (out response.WeightResponse, httpCode int, code string, err error) {
	var model sql.TxWeightRecord

	inJSON, _ := json.Marshal(in)
	json.Unmarshal(inJSON, &model)

	model.WeightRecordId = nil

	res, err := uc.weightUC.UpdateWeight(model, "weight_record_id = ?", in.WeightRecordId)

	// convert map to json
	jsonString, _ := json.Marshal(res)
	json.Unmarshal(jsonString, &out)
	return
}
