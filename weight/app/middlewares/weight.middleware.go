package middlewares

import (
	configs "project/app/configs"
	structs "project/app/structs"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// AddWeight function
func AddWeight(c *gin.Context) {
	configs.Block{
		Try: func() {
			var request structs.AddWeight

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
			configs.FancyHandleError(e)
		},
	}.Do()
}

// UpdateWeight function
func UpdateWeight(c *gin.Context) {
	configs.Block{
		Try: func() {
			var request structs.UpdateWeight

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
			configs.FancyHandleError(e)
		},
	}.Do()
}

// GetWeight function
func GetWeight(c *gin.Context) {
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
			configs.FancyHandleError(e)
		},
	}.Do()
}

// GetWeightById function
func GetWeightById(c *gin.Context) {
	configs.Block{
		Try: func() {
			var request structs.GetWeight

			weightRecordId := c.Param("weightRecordId")
			request.WeightRecordId = weightRecordId

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
			configs.FancyHandleError(e)
		},
	}.Do()
}
