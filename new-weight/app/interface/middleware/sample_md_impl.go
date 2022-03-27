package middleware

import (
	configs "project-name/app/shared/utils"

	structs "project-name/app/shared/structs"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func (m *md) GetWeightById(c *gin.Context) {
	configs.Block{
		Try: func() {
			var request structs.GetWeight

			weightRecordId := c.Param("weightRecordId")
			request.WeightRecordId = &weightRecordId

			if err := v.Struct(request); err != nil {
				configs.Throw(err)
			}

			c.Next()
		},
		Catch: func(e error) {
			c.AbortWithStatusJSON(400, gin.H{
				"status":  "FAILED",
				"message": e.Error(),
				"data":    nil,
			})
		},
	}.Do()
}

func (m *md) GetWeight(c *gin.Context) {
	configs.Block{
		Try: func() {
			var request structs.GetWeight

			if err := c.ShouldBindQuery(&request); err != nil {
				configs.Throw(err)
			}

			if err := v.Struct(request); err != nil {
				configs.Throw(err)
			}

			c.Next()
		},
		Catch: func(e error) {
			c.AbortWithStatusJSON(400, gin.H{
				"status":  "FAILED",
				"message": e.Error(),
				"data":    nil,
			})
		},
	}.Do()
}

func (m *md) AddWeight(c *gin.Context) {
	configs.Block{
		Try: func() {
			var request structs.WeightData

			request.WeightRecordId = nil

			if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
				configs.Throw(err)
			}

			if err := v.Struct(request); err != nil {
				configs.Throw(err)
			}

			c.Next()
		},
		Catch: func(e error) {
			c.AbortWithStatusJSON(400, gin.H{
				"status":  "FAILED",
				"message": e.Error(),
				"data":    nil,
			})
		},
	}.Do()
}

func (m *md) UpdateWeight(c *gin.Context) {
	configs.Block{
		Try: func() {
			var request structs.WeightData

			weightRecordId := c.Param("weightRecordId")
			request.WeightRecordId = &weightRecordId

			if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
				configs.Throw(err)
			}

			if err := v.Struct(request); err != nil {
				configs.Throw(err)
			}

			c.Next()
		},
		Catch: func(e error) {
			c.AbortWithStatusJSON(400, gin.H{
				"status":  "FAILED",
				"message": e.Error(),
				"data":    nil,
			})
		},
	}.Do()
}
