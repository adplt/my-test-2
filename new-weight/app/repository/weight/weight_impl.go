package weight

import (
	"project-name/app/adapter/database"
	"project-name/app/models/sql"
	"time"
)

type weightRepo struct {
	mysql *database.MySQL
}

func NewWeightRepo(mysql *database.MySQL) RepositoryWeight {
	return &weightRepo{
		mysql: mysql,
	}
}

func (s *weightRepo) FindAll(in string, limit, offset int, query string, where ...interface{}) (out []sql.TxWeightRecord, err error) {
	err = s.mysql.Find(&out).Limit(limit).Offset(offset).Error
	return
}

func (s *weightRepo) FindById(in string) (out sql.TxWeightRecord, err error) {
	err = s.mysql.Where("weight_record_id = ?", in).Find(&out).Limit(1).Error
	return
}

func (s *weightRepo) SaveWeight(in sql.TxWeightRecord) (out sql.TxWeightRecord, err error) {
	err = s.mysql.Omit("modified_date").Save(&in).Error
	return
}

func (s *weightRepo) UpdateWeight(in sql.TxWeightRecord) (out sql.TxWeightRecord, err error) {
	in.ModifiedDate = time.Now().Format("2006-01-02 15:04:05")
	in.ModifiedBy = "Backend"
	in.ModifiedFrom = "Backend"
	if err := s.mysql.Update(&in).Error; err != nil {
		return in, err
	}
	return
}
