package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"project-name/app/shared/config"
	"time"
)

type SQLServer struct {
	*gorm.DB
}

func SetupSqlServer() *SQLServer {
	fmt.Println("Try Setup MySQL ...")
	url := config.DB.SqlServer.DataSourceName
	dialect := config.DB.SqlServer.DriverName

	db, err := gorm.Open(dialect, url)
	if err != nil {
		panic(err)
	}

	// Enable Logger, show detailed log
	// db.LogMode(true)

	db.DB().SetConnMaxLifetime(5 * time.Minute)
	return &SQLServer{db}
}
