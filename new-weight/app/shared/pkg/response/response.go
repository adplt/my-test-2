package response

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ServerResponse godoc.
type ServerResponse interface {
	BadRequest(ctx *gin.Context)
	DataNotFound(ctx *gin.Context)
	InternalServerError(ctx *gin.Context)
	Error(ctx *gin.Context, httpStatusCode int, code string)
	SuccessOK(ctx *gin.Context, data interface{})
	Success(ctx *gin.Context, code, problemOwner string, data interface{})
	SuccessCode(ctx *gin.Context, code string)
	ResponseWithCode(ctx *gin.Context, httpStatusCode int, status bool, code, problemOwner string, data interface{})
}

type serverResponse struct {
	db *sql.DB
}

// New godoc.
func New(db *sql.DB) ServerResponse {
	return &serverResponse{db}
}

// BadRequest godoc.
func (r serverResponse) BadRequest(ctx *gin.Context) {
	r.Error(ctx, http.StatusBadRequest, StatusDataNotValid)
}

// DataNotFound godoc.
func (r serverResponse) DataNotFound(ctx *gin.Context) {
	r.Error(ctx, http.StatusNotFound, StatusDataNotFound)
}

// InternalServerError godoc.
func (r serverResponse) InternalServerError(ctx *gin.Context) {
	r.Error(ctx, http.StatusInternalServerError, StatusInternalServerError)
}

// Error godoc.
func (r serverResponse) Error(ctx *gin.Context, httpStatusCode int, code string) {
	r.ErrorPO(ctx, httpStatusCode, code, StatusOwner)
}

func (r serverResponse) ErrorPO(ctx *gin.Context, httpStatusCode int, code, problemOwner string) {
	if code == "" {
		code = StatusInternalServerError
	}
	if problemOwner == "" {
		problemOwner = StatusOwner
	}

	r.basic(ctx, httpStatusCode, false, code, problemOwner, nil)
}

// Success godoc
func (r serverResponse) Success(ctx *gin.Context, code, problemOwner string, data interface{}) {
	if code == "" {
		code = StatusOK
	}
	if problemOwner == "" {
		problemOwner = StatusOwner
	}

	r.basic(ctx, http.StatusOK, true, code, problemOwner, data)
}

func (r serverResponse) SuccessCode(ctx *gin.Context, code string) {
	r.Success(ctx, code, StatusOwner, nil)
}

// SuccessOK godoc.
func (r serverResponse) SuccessOK(ctx *gin.Context, data interface{}) {
	r.Success(ctx, StatusOK, StatusOwner, data)
}

// ResponseWithCode godoc
func (r serverResponse) ResponseWithCode(ctx *gin.Context, httpStatusCode int, status bool, code, problemOwner string, data interface{}) {
	if code == "" {
		code = StatusOK
	}
	if problemOwner == "" {
		problemOwner = StatusOwner
	}

	r.basic(ctx, httpStatusCode, status, code, problemOwner, data)
}

func (r serverResponse) basic(ctx *gin.Context, httpStatusCode int, status bool, code, problemOwner string, data interface{}) {
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

	Basic(status, httpStatusCode, code, msgID+"\n"+msgEN, data).JSON(ctx)
}
