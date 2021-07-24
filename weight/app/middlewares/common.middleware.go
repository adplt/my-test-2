package middlewares

import (
	configs "project/app/configs"
	"regexp"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

var (
	v          = validator.New()
	reNumber   = regexp.MustCompile("[^0-9]")
	reDate     = regexp.MustCompile("([12]\\d{3}-(0[1-9]|1[0-2])-(0[1-9]|[12]\\d|3[01]))")
	reAlphanum = regexp.MustCompile("[^0-9a-zA-Z ]")
)

// CORS function
func CORS() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		context.Writer.Header().Set("Access-Control-Max-Age", "86400")
		context.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		context.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		context.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if context.Request.Method == "OPTIONS" {
			context.AbortWithStatus(200)
		} else {
			context.Next()
		}
	}
}

// JSONLogMiddleware function
func JSONLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		configs.Block{
			Try: func() {
			},
			Catch: func(e error) {
				configs.FancyHandleError(e)
			},
		}.Do()
	}
}
