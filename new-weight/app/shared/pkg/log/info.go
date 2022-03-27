package log

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

// InfoMsgs godoc.
type InfoMsgs []interface{}

// InfoKey godoc.
const InfoKey = "LoggerInfoKey"

// Info godoc.
func Info(ctx *gin.Context, in interface{}) {
	msg := ParseMessage(fmt.Sprintf("%v %v", time.Now().Format("2006-01-02 15:04:05"), in))

	if ctx != nil {
		infoMsgs := parseToInfoMsgs(ctx)
		infoMsgs = append(infoMsgs, parseToInterface(msg))
		ctx.Set(InfoKey, infoMsgs)
	} else {
		// send traffic log to log
	}
}

// InfoG godoc.
func InfoG(in interface{}) {
	// notice that we're using 1, so it will actually log the where
	// the error happened, 0 = this function, we don't want that.
	pc, fn, line, _ := runtime.Caller(1)
	msg := ParseMessage(fmt.Sprintf("%s[%s:%d] %v", runtime.FuncForPC(pc).Name(), fn, line, in))

	// send traffic log to log

	log.Printf("[info] in %v", msg)
}

// InfoGetFromContext godoc.
func InfoGetFromContext(ctx *gin.Context) InfoMsgs {
	infoMsgs := parseToInfoMsgs(ctx)
	return infoMsgs
}

func parseToInterface(i interface{}) interface{} {
	return i
}

func parseToInfoMsgs(ctx *gin.Context) (infoMsgs InfoMsgs) {
	infoMsgs = make(InfoMsgs, 0)
	ctxValue := ctx.Value(InfoKey)
	if ctxValue != nil {
		infoMsgs, _ = ctxValue.(InfoMsgs)
	}
	return
}
