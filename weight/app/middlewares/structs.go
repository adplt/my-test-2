package middlewares

import (
	"bytes"

	"github.com/gin-gonic/gin"
)

// BodyLogWriter struct
type BodyLogWriter struct {
	body *bytes.Buffer
	gin.ResponseWriter
}

func (w BodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// WriteString function
func (w BodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}
