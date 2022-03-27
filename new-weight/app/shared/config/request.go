package config

import (
	"time"

	"github.com/gomodul/envy"
	"github.com/spf13/cast"
)

type request struct {
	Timeout time.Duration
}

var Request = request{
	Timeout: cast.ToDuration(envy.Get("REQUEST_TIMEOUT", "30")) * time.Second,
}
