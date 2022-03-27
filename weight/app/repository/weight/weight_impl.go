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

func (s *weightRepo) FindAll(limit, offset int, query string, where ...interface{}) (out []sql.TxWeightRecord, err error) {
	err = s.mysql.Limit(limit).Offset(offset).Where(query, where).Find(&out).Error
	return
}

func (s *weightRepo) FindById(in string) (out sql.TxWeightRecord, err error) {
	err = s.mysql.Where("weight_record_id = ?", in).Find(&out).Limit(1).Error
	return
}

func (s *weightRepo) SaveWeight(in sql.TxWeightRecord) (out sql.TxWeightRecord, err error) {
	in.CreatedDate = time.Now().Format("2006-01-02 15:04:05")
	in.CreatedBy = "Backend"
	in.CreatedFrom = "Backend"
	err = s.mysql.Omit("modified_date").Save(&in).Error
	return in, nil
}

func (s *weightRepo) UpdateWeight(in sql.TxWeightRecord, query string, where ...interface{}) (out sql.TxWeightRecord, err error) {
	in.ModifiedDate = time.Now().Format("2006-01-02 15:04:05")
	in.ModifiedBy = "Backend"
	in.ModifiedFrom = "Backend"
	if err := s.mysql.Table("tx_weight_record").Where(query, where).Update(&in).Error; err != nil {
		return sql.TxWeightRecord{}, err
	}
	return in, nil
}
