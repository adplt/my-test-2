package main

import (
	"net/http"
	"os"
	configs "project/app/configs"
	variables "project/app/variables"
	"project/docs"
	"strings"

	"github.com/gin-gonic/gin"

	controllers "project/app/controllers"
	middlewares "project/app/middlewares"

	_ "project/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Weight
// @version 1.0
// @description Example API for Weight
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	// gin.SetMode(gin.ReleaseMode)
	docs.SwaggerInfo.Title = "Weight"
	docs.SwaggerInfo.Description = "Example API for Weight"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = configs.Env.SwaggerURL

	r := gin.New()
	r.LoadHTMLGlob("app/views/*")
	os.Setenv("ENV", "prod")
	r.Use(middlewares.CORS())
	r.Use(middlewares.JSONLogMiddleware())
	r.GET("/", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"title": "Home Page",
			},
		)

	})
	if strings.ToUpper(configs.Env.SwaggerAllow) == "TRUE" {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	v1 := r.Group(variables.API_VERSION)
	{
		v1weight := v1.Group(variables.WEIGHT)
		v1weight.GET("/", middlewares.GetWeight, controllers.GetWeight)
		v1weight.POST("/", middlewares.AddWeight, controllers.AddWeight)
		v1weight.PUT("/:weightRecordId", middlewares.UpdateWeight, controllers.UpdateWeight)
	}

	r.Run(":" + configs.Env.AppPort)
}
