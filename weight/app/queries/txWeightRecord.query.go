package queries

import (
	"errors"
	configs "project/app/configs"
	models "project/app/models"
	structs "project/app/structs"
	utils "project/app/utils"
	variables "project/app/variables"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func TxWeightRecordSave(payload *models.TxWeightRecord) *models.TxWeightRecord {
	db, err := configs.DbConnection1()
	if err != nil {
		configs.Throw(err)
	}
	res := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&payload)
	if res.Error != nil {
		configs.Throw(res.Error)
	}
	sqlDB, err := db.DB()
	if err != nil {
		configs.Throw(err)
	}
	defer sqlDB.Close()
	return payload
}

func TxWeightRecordFindByPk(attributes []string, id string) *models.TxWeightRecord {
	result := &models.TxWeightRecord{}
	db, err := configs.DbConnection1()
	if err != nil {
		configs.Throw(err)
	}
	db = utils.SetAttribute(db, attributes...)
	if res := db.First(&result, "weight_record_id = ? AND status_id = ?", id, variables.ACTIVE_STATUS); res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return &models.TxWeightRecord{}
		} else {
			configs.Throw(res.Error)
		}
	}
	sqlDB, err := db.DB()
	if err != nil {
		configs.Throw(err)
	}
	defer sqlDB.Close()
	return result
}

func TxWeightRecordFindOne(attributes []string, orders []string, query string, payload ...interface{}) *models.TxWeightRecord {
	var result *models.TxWeightRecord
	db, err := configs.DbConnection1()
	if err != nil {
		configs.Throw(err)
	}
	db = utils.SetAttribute(db, attributes...)
	db = utils.SetOrder(db, orders)
	if res := db.Where(query, payload...).First(&result); res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return &models.TxWeightRecord{}
		} else {
			configs.Throw(res.Error)
		}
	}
	sqlDB, err := db.DB()
	if err != nil {
		configs.Throw(err)
	}
	defer sqlDB.Close()
	return result
}

func TxWeightRecordFind(attributes []string, orders []string, query string, payload ...interface{}) []*structs.TxWeightRecord {
	result := []*structs.TxWeightRecord{}
	db, err := configs.DbConnection1()
	if err != nil {
		configs.Throw(err)
	}
	db = utils.SetAttribute(db, attributes...)
	db = utils.SetOrder(db, orders)
	if res := db.Where(query, payload...).Find(&result); res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return []*structs.TxWeightRecord{}
		} else {
			configs.Throw(res.Error)
		}
	}
	sqlDB, err := db.DB()
	if err != nil {
		configs.Throw(err)
	}
	defer sqlDB.Close()
	return result
}

func TxWeightRecordUpdate(payload *models.TxWeightRecord, query string, where ...interface{}) *models.TxWeightRecord {
	db, err := configs.DbConnection1()
	if err != nil {
		configs.Throw(err)
	}
	res := db.Where(query, where...).Updates(&payload)
	if res.Error != nil {
		configs.Throw(res.Error)
	}
	sqlDB, err := db.DB()
	if err != nil {
		configs.Throw(err)
	}
	defer sqlDB.Close()
	return payload
}
