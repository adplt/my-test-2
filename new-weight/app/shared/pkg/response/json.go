package response

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// JSON struct
type JSON struct {
	HTTPStatusCode int         `json:"-"`
	Success        bool        `json:"success"`
	MessageCode    string      `json:"messagecode"`
	Message        string      `json:"message"`
	Results        interface{} `json:"results"`
}

// Basic godoc.
func Basic(status bool, httpStatusCode int, code, message string, data interface{}) *JSON {
	return &JSON{
		Success:        status,
		HTTPStatusCode: httpStatusCode,
		MessageCode:    code,
		Message:        message,
		Results:        data,
	}
}

// JSON godoc.
func (res JSON) JSON(c *gin.Context, data ...interface{}) {
	if !res.Success && data != nil && data[0] != nil {
		res.Results = data[0]
	}
	c.JSON(res.HTTPStatusCode, gin.H{})
}

// Error godoc.
func (res *JSON) Error() error {
	return errors.New(res.Message)
}

// JSONBasic godoc.
func (res JSON) JSONBasic(c *gin.Context, data ...interface{}) {
	if !res.Success && data != nil && data[0] != nil {
		res.Results = data[0]
	}

	c.AbortWithStatusJSON(res.HTTPStatusCode, res)
}
