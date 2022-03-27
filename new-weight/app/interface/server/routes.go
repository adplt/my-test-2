package server

import (
	middlewares "project-name/app/interface/middleware"
	"project-name/app/transport"

	"github.com/gin-gonic/gin"
	"github.com/gomodul/envy"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func setupRouter(transport *transport.Tp, middleware middlewares.Middleware, app *gin.Engine) {
	// allow swagger
	if envy.Get("ALLOW_SWAGGER", "false") == "true" {
		app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// this router group for readiness and health check
	ping := app.Group("/ping")
	ping.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	v1 := app.Group("/v1")
	v1.GET("/:weightRecordId/", middleware.GetWeightById, transport.SampleTransport.GetWeightById)
	v1.GET("/", middleware.GetWeight, transport.SampleTransport.GetWeight)
	v1.POST("/", middleware.AddWeight, transport.SampleTransport.AddWeight)
	v1.PUT("/:weightRecordId/", middleware.UpdateWeight, transport.SampleTransport.UpdateWeight)
}
