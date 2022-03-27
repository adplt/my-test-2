package database

import (
	"fmt"
	"project-name/app/shared/config"
	"time"

	"github.com/jinzhu/gorm"
)

type MySQL struct {
	*gorm.DB
}

func SetupMySQL() *MySQL {
	fmt.Println("Try Setup MySQL ...")
	url := config.DB.MySQL.DataSourceName
	dialect := config.DB.MySQL.DriverName

	db, err := gorm.Open(dialect, url)
	if err != nil {
		panic(err)
	}

	// Enable Logger, show detailed log
	db.LogMode(true)

	db.DB().SetConnMaxLifetime(5 * time.Minute)

	return &MySQL{db}
}
