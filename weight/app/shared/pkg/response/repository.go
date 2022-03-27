package response

import (
	"context"
	"database/sql"
)

const readErrorDetailMessage = `SELECT wording_i, wording_e from mb_error_details where error_code=? and problem_owner=?`

// ReadErrorDetailMessageParams godoc.
type ReadErrorDetailMessageParams struct {
	ErrorCode    string `db:"error_code" json:"error_code"`
	ProblemOwner string `db:"problem_owner" json:"problem_owner"`
}

// ReadErrorDetailMessage godoc.
func ReadErrorDetailMessage(ctx context.Context, db *sql.DB, arg ReadErrorDetailMessageParams) *sql.Row {
	return db.QueryRowContext(ctx, readErrorDetailMessage, arg.ErrorCode, arg.ProblemOwner)
}
