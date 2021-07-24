package tests

import (
	"errors"
	"fmt"
	configs "project/app/configs"
	"testing"

	structs "project/app/structs"
	testutils "project/app/tests/utils"

	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/imdario/mergo"
)

/* e2e test */

var (
	weightRecordId string
	commonReq      = map[string]interface{}{
		"userId":   "Postman",
		"userName": "Postman",
		"source":   "Postman",
	}
)

func TestAddWeight(t *testing.T) {
	configs.Block{
		Try: func() {
			var payload structs.ResponseCommonObject
			request := map[string]interface{}{
				"date": "2021-07-24",
				"max":  55,
				"min":  50,
			}
			mergo.Merge(&request, commonReq)
			res, errRes := testutils.AddWeightAPI(request)
			if errRes != nil {
				fmt.Println("err: ", string(errRes.Data.([]uint8)))
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
			weightRecordId = payload.Data["weight_record_id"].(string)
		},
		Catch: func(e error) {
			t.Fatalf("Fatal: %v", e)
		},
	}.Do()
}

func TestUpdateWeight(t *testing.T) {
	configs.Block{
		Try: func() {
			var payload structs.ResponseCommonObject
			request := map[string]interface{}{
				"date": "2021-07-24",
				"max":  55,
				"min":  50,
			}
			mergo.Merge(&request, commonReq)
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
		},
		Catch: func(e error) {
			t.Fatalf("Fatal: %v", e)
		},
	}.Do()
}

func TestGetWeight(t *testing.T) {
	configs.Block{
		Try: func() {
			var payload structs.ResponseCommonArray
			res, errRes := testutils.GetWeightAPI("?date=2021-07-24")
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
		},
		Catch: func(e error) {
			t.Fatalf("Fatal: %v", e)
		},
	}.Do()
}
