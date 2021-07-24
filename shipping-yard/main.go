package main

import (
	"net/http"
	"os"
	configs "project/app/configs"
	"project/docs"
	"strings"

	"github.com/gin-gonic/gin"

	middlewares "project/app/middlewares"

	_ "project/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Shipping Yard
// @version 1.0
// @description Example API for Shipping Yard
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	// gin.SetMode(gin.ReleaseMode)
	docs.SwaggerInfo.Title = "Shipping Yard"
	docs.SwaggerInfo.Description = "Example API for Shipping Yard"
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
	r.Run(":" + configs.Env.AppPort)
}
