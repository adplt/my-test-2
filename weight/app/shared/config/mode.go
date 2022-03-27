package config

import "github.com/gomodul/envy"

type appenv struct {
	APPENV string
}

var AppEnv = appenv{
	APPENV: envy.Get("APP_ENV", "development"),
}
