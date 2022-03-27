package log

import (
	"bytes"
	"encoding/json"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type middleware struct {
	// TimeStamp shows the time after the server returns a response.
	TimeStamp string `json:"timestamp"`
	// StatusCode is HTTP response code.
	StatusCode int `json:"statuscode"`
	// Latency is how much time the server cost to process a certain request.
	Latency float64 `json:"latency"`
	// ClientIP equals Context's ClientIP method.
	ClientIP string `json:"clientip"`
	// Method is the HTTP method given to the request.
	Method string `json:"method"`
	// Path is a path the client requests.
	Path string `json:"path"`
	// BodySize is the size of the Response Body
	BodySize int `json:"bodysize"`
	// Info godoc.
	Info InfoMsgs `json:"info"`
	// Error godoc.
	Error []string `json:"error"`
	// Request body
	Request map[string]interface{} `json:"request"`
	// Response body
	Response interface{} `json:"response"`
}

// Middleware for middleware with encyption godoc.
func MiddlewareSecure(ctx *gin.Context) {

}

// MiddlewareBasic is for middleware without encyption.
func MiddlewareBasic(ctx *gin.Context) {

}

func printLog(logJSON *middleware) {
	logByte, _ := json.MarshalIndent(logJSON, "", " ")
	logString := string(logByte)

	// send traffic log to log

	log.SetFlags(0)
	log.Print(logString)
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// Write godoc.
func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func parseBody(body []byte) (res map[string]interface{}) {
	_ = json.Unmarshal(body, &res)
	return
}

func getDurationInMilliseconds(start time.Time) float64 {
	end := time.Now()
	duration := end.Sub(start)
	milliseconds := float64(duration) / float64(time.Millisecond)
	rounded := float64(int(milliseconds*100+.5)) / 100
	return rounded
}
