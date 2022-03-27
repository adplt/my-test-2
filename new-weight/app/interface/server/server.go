package server

import (
	"fmt"
	"project-name/app/interface/container"
	"project-name/app/interface/middleware"
	"project-name/app/shared/config"
	"project-name/app/transport"

	"github.com/gin-gonic/gin"
)

func StartServer(container container.Container) {
	// init gin
	app := gin.New()

	// setup Middle ware
	middle := middleware.NewMiddleware(&container)

	// setup transport and router (parent)
	transport := transport.SetupTransport(&container)
	setupRouter(transport, middle, app)

	// run gin apps
	fmt.Println(app.Run(config.Server.PORTHTTP))
}
