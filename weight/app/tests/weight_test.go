package tests

import (
	"errors"
	"fmt"
	configs "project/app/configs"
	"testing"

	structs "project/app/structs"
	testutils "project/app/tests/utils"

	variables "project/app/variables"

	models "project/app/models"

	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/imdario/mergo"
)

/* e2e testing */

var (
	weightRecordId string
	userId         = "Postman"
	userName       = "Postman"
	source         = "Postman"
	commonReq      = map[string]interface{}{
		"userId":   userId,
		"userName": userName,
		"source":   source,
	}
)

func TestAddWeight(t *testing.T) {
	configs.Block{
		Try: func() {
			/* Set the input */
			var (
				payload   structs.ResponseCommonObject
				inputDate = "2021-07-21"
				inputMax  = 55
				inputMin  = 50
			)
			request := map[string]interface{}{
				"date": inputDate,
				"max":  inputMax,
				"min":  inputMin,
			}
			mergo.Merge(&request, commonReq)
			/* Call the API & Check API Status */
			res, errRes := testutils.AddWeightAPI(request)
			if errRes != nil {
				configs.Throw(errors.New("Main API Error: " + errRes.Message))
			}
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				configs.Throw(err)
			}
			reader := strings.NewReader(string(body))
			_ = json.NewDecoder(reader).Decode(&payload)
			if payload.Data == nil {
				configs.Throw(errors.New("Main API Response Failed: " + fmt.Sprintf("%v", payload.Message) + " - " + fmt.Sprintf("%v", payload.Error)))
			}
			/* Check the payload */
			var (
				weight *models.TxWeightRecord
			)
			dataJson, err := json.Marshal(payload.Data)
			if err != nil {
				configs.Throw(err)
			}
			err = json.Unmarshal(dataJson, &weight)
			if err != nil {
				configs.Throw(err)
			}
			weightRecordId = weight.WeightRecordId.String()
			if weightRecordId == "" || weightRecordId == variables.EMPTY_GUID {
				configs.Throw(errors.New("Weight Record ID not found"))
			}
			if weight.Date != inputDate {
				configs.Throw(errors.New("Wrong date response"))
			}
			if weight.Max != inputMax {
				configs.Throw(errors.New("Wrong max response"))
			}
			if weight.Min != inputMin {
				configs.Throw(errors.New("Wrong min response"))
			}
			if weight.StatusId != variables.ACTIVE_STATUS {
				configs.Throw(errors.New("Status is not active"))
			}
			if weight.CreatedBy != userId {
				configs.Throw(errors.New("Wrong created by"))
			}
			if weight.CreatedDate == "" {
				configs.Throw(errors.New("Wrong created name"))
			}
			if weight.CreatedName != userName {
				configs.Throw(errors.New("Wrong created name"))
			}
			if weight.CreatedFrom != source {
				configs.Throw(errors.New("Wrong created from"))
			}
			if weight.ModifiedBy != userId {
				configs.Throw(errors.New("Wrong modified by"))
			}
			if weight.ModifiedDate == "" {
				configs.Throw(errors.New("Wrong modified name"))
			}
			if weight.ModifiedName != userName {
				configs.Throw(errors.New("Wrong modified name"))
			}
			if weight.ModifiedFrom != source {
				configs.Throw(errors.New("Wrong modified from"))
			}
		},
		Catch: func(e error) {
			t.Fatalf("Fatal: %v", e)
		},
	}.Do()
}

func TestUpdateWeight(t *testing.T) {
	configs.Block{
		Try: func() {
			/* Set the input */
			var (
				payload   structs.ResponseCommonObject
				inputDate = "2021-07-24"
			)
			request := map[string]interface{}{
				"date": inputDate,
			}
			mergo.Merge(&request, commonReq)
			/* Call the API & Check API Status */
			res, errRes := testutils.UpdateWeightAPI(weightRecordId, request)
			if errRes != nil {
				configs.Throw(errors.New("Main API Error: " + errRes.Message))
			}
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				configs.Throw(err)
			}
			reader := strings.NewReader(string(body))
			_ = json.NewDecoder(reader).Decode(&payload)
			if payload.Data == nil {
				configs.Throw(errors.New("Main API Response Failed: " + fmt.Sprintf("%v", payload.Message) + " - " + fmt.Sprintf("%v", payload.Error)))
			}
			/* Check the payload */
			var (
				weight *models.TxWeightRecord
			)
			dataJson, err := json.Marshal(payload.Data)
			if err != nil {
				configs.Throw(err)
			}
			err = json.Unmarshal(dataJson, &weight)
			if err != nil {
				configs.Throw(err)
			}
			if weight.WeightRecordId.String() != weightRecordId {
				configs.Throw(errors.New("Wrong weight record ID"))
			}
			if weight.Date != inputDate {
				configs.Throw(errors.New("Wrong date response"))
			}
			if weight.ModifiedBy != userId {
				configs.Throw(errors.New("Wrong modified by"))
			}
			if weight.ModifiedDate == "" {
				configs.Throw(errors.New("Wrong modified name"))
			}
			if weight.ModifiedName != userName {
				configs.Throw(errors.New("Wrong modified name"))
			}
			if weight.ModifiedFrom != source {
				configs.Throw(errors.New("Wrong modified from"))
			}
		},
		Catch: func(e error) {
			t.Fatalf("Fatal: %v", e)
		},
	}.Do()
}

func TestGetWeight(t *testing.T) {
	configs.Block{
		Try: func() {
			/* Set the input */
			var (
				payload   structs.ResponseCommonArray
				inputDate = "2021-07-24"
			)
			/* Call the API & Check API Status */
			res, errRes := testutils.GetWeightAPI("?date=" + inputDate)
			if errRes != nil {
				configs.Throw(errors.New("Main API Error: " + errRes.Message))
			}
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				configs.Throw(err)
			}
			reader := strings.NewReader(string(body))
			_ = json.NewDecoder(reader).Decode(&payload)
			if len(payload.Data) == 0 {
				configs.Throw(errors.New("Main API Response Failed: " + fmt.Sprintf("%v", payload.Message) + " - " + fmt.Sprintf("%v", payload.Error)))
			}
			/* Check the payload */
			var (
				weight []*models.TxWeightRecord
			)
			dataJson, err := json.Marshal(payload.Data)
			if err != nil {
				configs.Throw(err)
			}
			err = json.Unmarshal(dataJson, &weight)
			if err != nil {
				configs.Throw(err)
			}
			_, err = json.Marshal(weight)
			if len(weight) == 0 {
				configs.Throw(errors.New("Response not found"))
			}
			if weight[0].WeightRecordId.String() == "" || weight[0].WeightRecordId.String() == variables.EMPTY_GUID {
				configs.Throw(errors.New("Wrong weight record ID"))
			}
			if weight[0].Date[:10] != inputDate {
				configs.Throw(errors.New("Wrong date response"))
			}
			if weight[0].StatusId != variables.ACTIVE_STATUS {
				configs.Throw(errors.New("Status is not active"))
			}
			if weight[0].CreatedBy != userId {
				configs.Throw(errors.New("Wrong created by"))
			}
			if weight[0].CreatedDate == "" {
				configs.Throw(errors.New("Wrong created name"))
			}
			if weight[0].CreatedName != userName {
				configs.Throw(errors.New("Wrong created name"))
			}
			if weight[0].CreatedFrom != source {
				configs.Throw(errors.New("Wrong created from"))
			}
			if weight[0].ModifiedBy != userId {
				configs.Throw(errors.New("Wrong modified by"))
			}
			if weight[0].ModifiedDate == "" {
				configs.Throw(errors.New("Wrong modified name"))
			}
			if weight[0].ModifiedName != userName {
				configs.Throw(errors.New("Wrong modified name"))
			}
			if weight[0].ModifiedFrom != source {
				configs.Throw(errors.New("Wrong modified from"))
			}
		},
		Catch: func(e error) {
			t.Fatalf("Fatal: %v", e)
		},
	}.Do()
}

func TestGetWeightById(t *testing.T) {
	configs.Block{
		Try: func() {
			/* Set the input */
			var (
				payload structs.ResponseCommonObject
			)
			/* Call the API & Check API Status */
			res, errRes := testutils.GetWeightByIdAPI(weightRecordId)
			if errRes != nil {
				configs.Throw(errors.New("Main API Error: " + errRes.Message))
			}
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				configs.Throw(err)
			}
			reader := strings.NewReader(string(body))
			_ = json.NewDecoder(reader).Decode(&payload)
			if payload.Data == nil {
				configs.Throw(errors.New("Main API Response Failed: " + fmt.Sprintf("%v", payload.Message) + " - " + fmt.Sprintf("%v", payload.Error)))
			}
			/* Check the payload */
			var (
				weight *models.TxWeightRecord
			)
			dataJson, err := json.Marshal(payload.Data)
			if err != nil {
				configs.Throw(err)
			}
			err = json.Unmarshal(dataJson, &weight)
			if err != nil {
				configs.Throw(err)
			}
			if weight.WeightRecordId.String() != weightRecordId {
				configs.Throw(errors.New("Wrong weight record ID"))
			}
			if weight.Date[:10] == "" {
				configs.Throw(errors.New("Wrong date response"))
			}
			if weight.Max == 0 {
				configs.Throw(errors.New("Wrong max response"))
			}
			if weight.Min == 0 {
				configs.Throw(errors.New("Wrong min response"))
			}
			if weight.StatusId != variables.ACTIVE_STATUS {
				configs.Throw(errors.New("Status is not active"))
			}
			if weight.CreatedBy != userId {
				configs.Throw(errors.New("Wrong created by"))
			}
			if weight.CreatedDate == "" {
				configs.Throw(errors.New("Wrong created name"))
			}
			if weight.CreatedName != userName {
				configs.Throw(errors.New("Wrong created name"))
			}
			if weight.CreatedFrom != source {
				configs.Throw(errors.New("Wrong created from"))
			}
			if weight.ModifiedBy != userId {
				configs.Throw(errors.New("Wrong modified by"))
			}
			if weight.ModifiedDate == "" {
				configs.Throw(errors.New("Wrong modified name"))
			}
			if weight.ModifiedName != userName {
				configs.Throw(errors.New("Wrong modified name"))
			}
			if weight.ModifiedFrom != source {
				configs.Throw(errors.New("Wrong modified from"))
			}
		},
		Catch: func(e error) {
			t.Fatalf("Fatal: %v", e)
		},
	}.Do()
}
