package testutil

import (
	"net"
	"net/http/httptest"
	"os"

	controllers "project/app/controllers"
	middlewares "project/app/middlewares"
	variables "project/app/variables"

	configs "project/app/configs"

	"github.com/gin-gonic/gin"
)

// ContentTest is
const ContentTest = "application/json; charset=utf-8"

// ServerTest function
func ServerTest() *httptest.Server {
	gin.SetMode(gin.TestMode)
	os.Setenv("ENV", "test")
	r := gin.New()
	r.Use(middlewares.CORS())
	r.Use(middlewares.JSONLogMiddleware())
	v1 := r.Group(variables.API_VERSION)
	{
		v1weight := v1.Group(variables.WEIGHT)
		v1weight.GET("/:weightRecordId", middlewares.GetWeightById, controllers.GetWeightById)
		v1weight.GET("/", middlewares.GetWeight, controllers.GetWeight)
		v1weight.POST("/", middlewares.AddWeight, controllers.AddWeight)
		v1weight.PUT("/:weightRecordId", middlewares.UpdateWeight, controllers.UpdateWeight)
	}
	ts := httptest.NewUnstartedServer(r)
	l, err := net.Listen("tcp", configs.Env.AppHost+":"+configs.Env.AppPort)
	if err != nil {
		panic(err)
	}

	ts.Listener.Close()
	ts.Listener = l
	ts.Start()

	return ts
}
