package weight

import (
	"project-name/app/models/sql"
)

type RepositoryWeight interface {
	FindAll(limit, offset int, query string, where ...interface{}) (out []sql.TxWeightRecord, err error)
	FindById(in string) (out sql.TxWeightRecord, err error)
	SaveWeight(in sql.TxWeightRecord) (out sql.TxWeightRecord, err error)
	UpdateWeight(in sql.TxWeightRecord) (out sql.TxWeightRecord, err error)
}
