package response

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

// BasicServerResponse godoc.
type BasicServerResponse interface {
	BadRequest(ctx *gin.Context)
	DataNotFound(ctx *gin.Context)
	InternalServerError(ctx *gin.Context)
	Error(ctx *gin.Context, httpStatusCode int, code string)
	SuccessOK(ctx *gin.Context, data interface{})
	Success(ctx *gin.Context, code, problemOwner string, data interface{})
	SuccessCode(ctx *gin.Context, code string)
	ResponseWithCode(ctx *gin.Context, httpStatusCode int, status bool, code, problemOwner string, data interface{})
}

type basicServerResponse struct {
	db *sql.DB
}

// NewBasic godoc.
func NewBasic(db *sql.DB) BasicServerResponse {
	return &basicServerResponse{db}
}

// BadRequest godoc.
func (r basicServerResponse) BadRequest(ctx *gin.Context) {
	r.Error(ctx, http.StatusBadRequest, StatusDataNotValid)
}

// DataNotFound godoc.
func (r basicServerResponse) DataNotFound(ctx *gin.Context) {
	r.Error(ctx, http.StatusNotFound, StatusDataNotFound)
}

// InternalServerError godoc.
func (r basicServerResponse) InternalServerError(ctx *gin.Context) {
	r.Error(ctx, http.StatusInternalServerError, StatusInternalServerError)
}

// Error godoc.
func (r basicServerResponse) Error(ctx *gin.Context, httpStatusCode int, code string) {
	r.ErrorPO(ctx, httpStatusCode, code, StatusOwner)
}

func (r basicServerResponse) ErrorPO(ctx *gin.Context, httpStatusCode int, code, problemOwner string) {
	if code == "" {
		code = StatusInternalServerError
	}
	if problemOwner == "" {
		problemOwner = StatusOwner
	}

	r.basic(ctx, httpStatusCode, false, code, problemOwner, nil)
}

// Success godoc
func (r basicServerResponse) Success(ctx *gin.Context, code, problemOwner string, data interface{}) {
	if code == "" {
		code = StatusOK
	}
	if problemOwner == "" {
		problemOwner = StatusOwner
	}

	r.basic(ctx, http.StatusOK, true, code, problemOwner, data)
}

func (r basicServerResponse) SuccessCode(ctx *gin.Context, code string) {
	r.Success(ctx, code, StatusOwner, nil)
}

// SuccessOK godoc.
func (r basicServerResponse) SuccessOK(ctx *gin.Context, data interface{}) {
	r.Success(ctx, StatusOK, StatusOwner, data)
}

// ResponseWithCode godoc
func (r basicServerResponse) ResponseWithCode(ctx *gin.Context, httpStatusCode int, status bool, code, problemOwner string, data interface{}) {
	if code == "" {
		code = StatusOK
	}
	if problemOwner == "" {
		problemOwner = StatusOwner
	}

	r.basic(ctx, httpStatusCode, status, code, problemOwner, data)
}

func (r basicServerResponse) basic(ctx *gin.Context, httpStatusCode int, status bool, code, problemOwner string, data interface{}) {
	var msgID, msgEN string

	row := ReadErrorDetailMessage(ctx, r.db, ReadErrorDetailMessageParams{
		ErrorCode:    code,
		ProblemOwner: problemOwner,
	})
	err := row.Scan(&msgID, &msgEN)
	if err != nil || msgID == "" || msgEN == "" {
		if status {
			msgID, msgEN = "Berhasil", "Success"
		} else {
			msgID, msgEN = "Error pada system", "System error"
		}
	}

	Basic(status, httpStatusCode, code, msgID+"\n"+msgEN, data).JSONBasic(ctx)
}
