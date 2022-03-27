package log

import (
	"errors"
	"fmt"
	"log"
	"runtime"

	"github.com/gin-gonic/gin"
)

// Error godoc.
func Error(ctx *gin.Context, err error) {
	if err != nil {
		// notice that we're using 1, so it will actually log the where
		// the error happened, 0 = this function, we don't want that.
		pc, fn, line, _ := runtime.Caller(1)
		msg := ParseMessage(fmt.Sprintf("%s[%s:%d] %v", runtime.FuncForPC(pc).Name(), fn, line, err.Error()))

		if ctx != nil {
			_ = ctx.Error(errors.New(msg))
		} else {
			// send traffic log to log

			log.Printf("[error] in %v", msg)
		}
	}
}
