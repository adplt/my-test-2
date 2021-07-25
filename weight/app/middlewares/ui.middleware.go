package middlewares

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	configs "project/app/configs"
	services "project/app/services"
	structs "project/app/structs"
	variables "project/app/variables"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/imdario/mergo"
)

// HomeUI function
func HomeUI() gin.HandlerFunc {
	return func(c *gin.Context) {
		configs.Block{
			Try: func() {
				httpRes, errSet := services.GetWeightAPI("")
				if errSet != nil {
					configs.Throw(errors.New(errSet.Message))
				}
				var payload map[string]interface{}
				body, err := ioutil.ReadAll(httpRes.Body)
				if err != nil {
					configs.Throw(err)
				}
				reader := strings.NewReader(string(body))
				err = json.NewDecoder(reader).Decode(&payload)
				if err != nil {
					configs.Throw(err)
				}
				payloadJson, err := json.Marshal(payload["data"])
				if err != nil {
					configs.Throw(err)
				}
				var (
					payloadStruct []*structs.TxWeightRecord
				)
				err = json.Unmarshal(payloadJson, &payloadStruct)
				if err != nil {
					configs.Throw(err)
				}
				c.HTML(
					http.StatusOK,
					"index.html",
					gin.H{
						"title": "Home Page",
						"data":  payloadStruct,
					},
				)
			},
			Catch: func(e error) {
				fmt.Println("Error HomeUI() function: ", e)
				configs.FancyHandleError(e)
			},
		}.Do()
	}
}

// SubmitWeightUI function
func SubmitWeightUI() gin.HandlerFunc {
	return func(c *gin.Context) {
		configs.Block{
			Try: func() {
				var (
					path string
				)
				weightRecordId := c.PostForm("weightRecordId")
				date := c.PostForm("date")
				max := c.PostForm("max")
				min := c.PostForm("min")

				commonReq := map[string]interface{}{
					"userId":   "Web",
					"userName": "Web",
					"source":   "Web",
				}

				if weightRecordId != "" && weightRecordId != variables.EMPTY_GUID {
					/* Update Weight */
					maxNum, err := strconv.Atoi(max)
					if err != nil {
						configs.Throw(err)
					}
					minNum, err := strconv.Atoi(min)
					if err != nil {
						configs.Throw(err)
					}
					request := map[string]interface{}{
						"date": date,
						"max":  maxNum,
						"min":  minNum,
					}
					mergo.Merge(&request, commonReq)
					httpRes, errSet := services.UpdateWeightAPI(weightRecordId, request)
					if errSet != nil {
						configs.Throw(errors.New(errSet.Message))
					}
					var payload map[string]interface{}
					body, err := ioutil.ReadAll(httpRes.Body)
					if err != nil {
						configs.Throw(err)
					}
					reader := strings.NewReader(string(body))
					err = json.NewDecoder(reader).Decode(&payload)
					if err != nil {
						configs.Throw(err)
					}
					payloadJson, err := json.Marshal(payload["data"])
					if err != nil {
						configs.Throw(err)
					}
					var (
						payloadStruct *structs.TxWeightRecord
					)
					err = json.Unmarshal(payloadJson, &payloadStruct)
					if err != nil {
						configs.Throw(err)
					}
					path = c.Request.URL.String()
				} else {
					/* Add Weight */
					maxNum, err := strconv.Atoi(max)
					if err != nil {
						configs.Throw(err)
					}
					minNum, err := strconv.Atoi(min)
					if err != nil {
						configs.Throw(err)
					}
					request := map[string]interface{}{
						"date": date,
						"max":  maxNum,
						"min":  minNum,
					}
					mergo.Merge(&request, commonReq)
					httpRes, errSet := services.AddWeightAPI(request)
					if errSet != nil {
						configs.Throw(errors.New(errSet.Message))
					}
					var payload map[string]interface{}
					body, err := ioutil.ReadAll(httpRes.Body)
					if err != nil {
						configs.Throw(err)
					}
					reader := strings.NewReader(string(body))
					err = json.NewDecoder(reader).Decode(&payload)
					if err != nil {
						configs.Throw(err)
					}
					payloadJson, err := json.Marshal(payload["data"])
					if err != nil {
						configs.Throw(err)
					}
					var (
						payloadStruct *structs.TxWeightRecord
					)
					err = json.Unmarshal(payloadJson, &payloadStruct)
					if err != nil {
						configs.Throw(err)
					}
					path = "/"
				}
				location := url.URL{Path: path}
				c.Redirect(http.StatusFound, location.RequestURI())
			},
			Catch: func(e error) {
				configs.FancyHandleError(e)
				fmt.Println("Error SubmitWeightUI() function: ", e)
			},
		}.Do()
	}
}

// FillWeightUI function
func FillWeightUI() gin.HandlerFunc {
	return func(c *gin.Context) {
		configs.Block{
			Try: func() {
				url := strings.Split(c.Request.URL.String(), "/")
				weightRecordId := url[len(url)-1]

				/* Inquiry by Detail */
				httpRes, errSet := services.GetWeightByIdAPI(weightRecordId)
				if errSet != nil {
					configs.Throw(errors.New(errSet.Message))
				}
				var payloadById map[string]interface{}
				body, err := ioutil.ReadAll(httpRes.Body)
				if err != nil {
					configs.Throw(err)
				}
				reader := strings.NewReader(string(body))
				err = json.NewDecoder(reader).Decode(&payloadById)
				if err != nil {
					configs.Throw(err)
				}
				payloadByIdJson, err := json.Marshal(payloadById["data"])
				if err != nil {
					configs.Throw(err)
				}
				var (
					payloadStructById *structs.TxWeightRecord
				)
				err = json.Unmarshal(payloadByIdJson, &payloadStructById)
				if err != nil {
					configs.Throw(err)
				}

				/* Get Weight List */
				httpRes, errSet = services.GetWeightAPI("")
				if errSet != nil {
					configs.Throw(errors.New(errSet.Message))
				}
				var payload map[string]interface{}
				body, err = ioutil.ReadAll(httpRes.Body)
				if err != nil {
					configs.Throw(err)
				}
				reader = strings.NewReader(string(body))
				err = json.NewDecoder(reader).Decode(&payload)
				if err != nil {
					configs.Throw(err)
				}
				payloadJson, err := json.Marshal(payload["data"])
				if err != nil {
					configs.Throw(err)
				}
				var (
					payloadStruct []*structs.TxWeightRecord
				)
				err = json.Unmarshal(payloadJson, &payloadStruct)
				if err != nil {
					configs.Throw(err)
				}

				c.HTML(
					http.StatusOK,
					"index.html",
					gin.H{
						"title":          "Home Page",
						"data":           payloadStruct,
						"weightRecordId": payloadStructById.WeightRecordId.String(),
						"date":           payloadStructById.Date[:10],
						"max":            payloadStructById.Max,
						"min":            payloadStructById.Min,
					},
				)
			},
			Catch: func(e error) {
				fmt.Println("Error FillWeightUI() function: ", e)
				configs.FancyHandleError(e)
			},
		}.Do()
	}
}
