package queries

import (
	"errors"
	configs "project/app/configs"
	models "project/app/models"
	utils "project/app/utils"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func TrLookupSave(payload *models.TrLookup) *models.TrLookup {
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

func TrLookupFindOne(attributes []string, orders []string, query string, payload ...interface{}) *models.TrLookup {
	var result *models.TrLookup
	db, err := configs.DbConnection1()
	if err != nil {
		configs.Throw(err)
	}
	db = utils.SetAttribute(db, attributes...)
	db = utils.SetOrder(db, orders)
	if res := db.Where(query, payload...).First(&result); res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return &models.TrLookup{}
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

func TrLookupFind(attributes []string, orders []string, query string, payload ...interface{}) []*models.TrLookup {
	result := []*models.TrLookup{}
	db, err := configs.DbConnection1()
	if err != nil {
		configs.Throw(err)
	}
	db = utils.SetAttribute(db, attributes...)
	db = utils.SetOrder(db, orders)
	if res := db.Where(query, payload...).Find(&result); res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return []*models.TrLookup{}
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

func TrLookupUpdate(payload *models.TrLookup, query string, where ...interface{}) *models.TrLookup {
	db, err := configs.DbConnection1()
	if err != nil {
		configs.Throw(err)
	}
	res := db.Where(query, where...).Updates(payload)
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
