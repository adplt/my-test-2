package middleware

import (
	"project-name/app/interface/container"
	"project-name/app/shared/pkg/response"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

var (
	v          = validator.New()
	reNumber   = regexp.MustCompile("[^0-9]")
	reDate     = regexp.MustCompile("([12]\\d{3}-(0[1-9]|1[0-2])-(0[1-9]|[12]\\d|3[01]))")
	reAlphanum = regexp.MustCompile("[^0-9a-zA-Z ]")
)

type md struct {
	responseWriter response.ServerResponse
}

// Middleware struct
type Middleware interface {
	GetWeightById(*gin.Context)
	GetWeight(*gin.Context)
	AddWeight(*gin.Context)
	UpdateWeight(ctx *gin.Context)
}

// NewMiddleware function
func NewMiddleware(container *container.Container) Middleware {
	return &md{
		responseWriter: container.ResponseParser,
	}
}
