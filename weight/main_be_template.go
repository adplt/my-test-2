package main

import (
	"project-name/app/interface/container"
	"project-name/app/interface/server"
)

func main() {
	server.StartServer(container.SetupContainer())
}
