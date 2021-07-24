package controllers

import (
	configs "project/app/configs"
	structs "project/app/structs"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// Login function
// @Summary API for Login
// @Description Login API
// @Param Body body structs.Login true "Body of Login"
// @Accept json
// @Produce json
// @Success 200 {object} structs.ResponseLogin
// @Failure 400 {object} structs.ResponseError400
// @Failure 500 {object} structs.ResponseError500
// @Router /api/v1/auth/login [post]
func SaveShip(c *gin.Context) {
	configs.Block{
		Try: func() {
			var request structs.Ship

			if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
				configs.Throw(err)
			}

			c.AbortWithStatusJSON(200, gin.H{
				"status":  "SUCCESS",
				"message": "Save ship successfully",
				"data":    request,
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
