package controllers

import (
	configs "project/app/configs"
	models "project/app/models"
	queries "project/app/queries"
	structs "project/app/structs"
	variables "project/app/variables"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
)

// AddWeight function
// @Summary API for Add Weight
// @Description Add Weight API
// @Param Body body structs.AddWeight true "Body of Weight"
// @Accept json
// @Produce json
// @Success 200 {object} structs.ResponseAddWeight
// @Failure 400 {object} structs.ResponseError400
// @Failure 500 {object} structs.ResponseError500
// @Router /api/v1/weights [post]
func AddWeight(c *gin.Context) {
	configs.Block{
		Try: func() {
			var request structs.AddWeight

			if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
				configs.Throw(err)
			}

			weightRecord := queries.TxWeightRecordSave(&models.TxWeightRecord{
				WeightRecordId: uuid.New(),
				Date:           request.Date,
				Max:            request.Max,
				Min:            request.Min,
				StatusId:       variables.ACTIVE_STATUS,
				CreatedBy:      request.UserId,
				CreatedDate:    time.Now().UTC().Format(time.RFC3339),
				CreatedName:    request.UserName,
				CreatedFrom:    request.Source,
				ModifiedBy:     request.UserId,
				ModifiedDate:   time.Now().UTC().Format(time.RFC3339),
				ModifiedName:   request.UserName,
				ModifiedFrom:   request.Source,
			})

			c.AbortWithStatusJSON(200, gin.H{
				"status":  "SUCCESS",
				"message": "Add weight successfully",
				"data":    weightRecord,
			})
		},
		Catch: func(e error) {
			c.AbortWithStatusJSON(500, gin.H{
				"status":  "FAILED",
				"message": e.Error(),
				"data":    nil,
			})
			configs.FancyHandleError(e)
		},
	}.Do()
}

// UpdateWeight function
// @Summary API for Update Weight
// @Description Update Weight API
// @Param weightRecordId path string true "Example: 419b68d8-fe6f-4025-b8e5-dcb572f74b27"
// @Param Body body structs.UpdateWeight true "Body of Weight"
// @Accept json
// @Produce json
// @Success 200 {object} structs.ResponseUpdateWeight
// @Failure 400 {object} structs.ResponseError400
// @Failure 500 {object} structs.ResponseError500
// @Router /api/v1/weights/{weightRecordId} [put]
func UpdateWeight(c *gin.Context) {
	configs.Block{
		Try: func() {
			var request structs.UpdateWeight

			if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
				configs.Throw(err)
			}

			weightRecordId := c.Param("weightRecordId")

			weightRecord := queries.TxWeightRecordUpdate(&models.TxWeightRecord{
				Date:         request.Date,
				Max:          request.Max,
				Min:          request.Min,
				ModifiedBy:   request.UserId,
				ModifiedDate: time.Now().UTC().Format(time.RFC3339),
				ModifiedName: request.UserName,
				ModifiedFrom: request.Source,
			}, "weight_record_id = ? AND status_id = ?", weightRecordId, variables.ACTIVE_STATUS)

			c.AbortWithStatusJSON(200, gin.H{
				"status":  "SUCCESS",
				"message": "Update weight successfully",
				"data":    weightRecord,
			})
		},
		Catch: func(e error) {
			c.AbortWithStatusJSON(500, gin.H{
				"status":  "FAILED",
				"message": e.Error(),
				"data":    nil,
			})
			configs.FancyHandleError(e)
		},
	}.Do()
}

// GetWeight function
// @Summary API for Add Weight
// @Description Add Weight API
// @Param date query string false "Example: 2021-04-01"
// @Accept json
// @Produce json
// @Success 200 {object} structs.ResponseGetWeight
// @Failure 400 {object} structs.ResponseError400
// @Failure 500 {object} structs.ResponseError500
// @Router /api/v1/weights [get]
func GetWeight(c *gin.Context) {
	configs.Block{
		Try: func() {
			var (
				request       structs.GetWeight
				weightRecords []*structs.TxWeightRecord
			)

			if err := c.ShouldBindQuery(&request); err != nil {
				configs.Throw(err)
			}

			if request.Date != "" {
				weightRecords = queries.TxWeightRecordFind([]string{
					"*",
					"max - min AS differences",
				}, []string{"max - min DESC"},
					"status_id = ? AND date = ?",
					variables.ACTIVE_STATUS,
					request.Date)
			} else {
				weightRecords = queries.TxWeightRecordFind([]string{
					"*",
					"max - min AS differences",
				}, []string{"max - min DESC"},
					"status_id = ?",
					variables.ACTIVE_STATUS)
			}

			c.AbortWithStatusJSON(200, gin.H{
				"status":  "SUCCESS",
				"message": "Get weight successfully",
				"data":    weightRecords,
			})
		},
		Catch: func(e error) {
			c.AbortWithStatusJSON(500, gin.H{
				"status":  "FAILED",
				"message": e.Error(),
				"data":    nil,
			})
			configs.FancyHandleError(e)
		},
	}.Do()
}
