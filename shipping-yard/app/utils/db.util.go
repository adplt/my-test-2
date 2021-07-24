package utils

import (
	"gorm.io/gorm"
)

func SetAttribute(db *gorm.DB, attribute ...string) *gorm.DB {
	return db.Select(attribute)
}

func SetOrder(db *gorm.DB, orders []string) *gorm.DB {
	for _, order := range orders {
		db.Order(order)
	}
	return db
}
