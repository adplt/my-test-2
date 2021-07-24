package testutil

import (
	"net"
	"net/http/httptest"
	"os"

	middlewares "project/app/middlewares"

	configs "project/app/configs"

	"github.com/gin-gonic/gin"
)

// ContentTest is
const ContentTest = "application/json; charset=utf-8"

// serverTest function
func serverTest() *httptest.Server {
	gin.SetMode(gin.TestMode)
	os.Setenv("ENV", "test")
	r := gin.New()
	r.Use(middlewares.CORS())
	r.Use(middlewares.JSONLogMiddleware())
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
