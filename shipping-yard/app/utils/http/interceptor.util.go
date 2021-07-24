package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Error400 function
func Error400(err error, data interface{}) *CommonJSON {
	return &CommonJSON{
		StatusCode: http.StatusBadRequest,
		Status:     "ERROR",
		Message:    err.Error(),
		Data:       data,
	}
}

// Error500 function
func Error500(err error, data interface{}) *CommonJSON {
	return &CommonJSON{
		StatusCode: http.StatusInternalServerError,
		Status:     "ERROR",
		Message:    err.Error(),
		Data:       data,
	}
}

// Error200 function
func Error200(json CommonJSON) *CommonJSON {
	json.StatusCode = http.StatusOK
	json.Status = "OK"
	return &json
}

// Success function
func Success(results interface{}) CommonJSON {
	return CommonJSON{
		StatusCode: http.StatusOK,
		Status:     "OK",
		Message:    "API Succesfully",
		Data:       results,
	}
}

// JSON function
func (res CommonJSON) JSON(c *gin.Context, data ...interface{}) {
	if res.StatusCode == 200 {
		c.AbortWithStatusJSON(res.StatusCode, res)
		return
	}
	c.AbortWithStatusJSON(res.StatusCode, res)
}
