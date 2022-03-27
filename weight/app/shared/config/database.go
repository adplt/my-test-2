package config

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodul/envy"
)

type databaseItem struct {
	DriverName     string
	DataSourceName string
}

type database struct {
	MySQL     databaseItem
	SqlServer databaseItem
}

var DB = database{
	MySQL: databaseItem{
		DriverName: envy.Get("MYSQL_DRIVER", "mysql"),
		DataSourceName: fmt.Sprintf(
			"%v:%v@tcp(%v:%v)/%v?parseTime=true&charset=utf8",
			envy.Get("MYSQL_USERNAME"),
			envy.Get("MYSQL_PASSWORD"),
			envy.Get("MYSQL_HOST"),
			envy.Get("MYSQL_PORT"),
			envy.Get("MYSQL_NAME"),
		),
	},
	SqlServer: databaseItem{
		DriverName: envy.Get("SQL_SERVER_DRIVER", "sqlserver"),
		DataSourceName: fmt.Sprintf(
			"server=%s;user id=%s;password=%s;port=%s;database=%s;",
			envy.Get("SQL_SERVER_HOST"),
			envy.Get("SQL_SERVER_USERNAME"),
			envy.Get("SQL_SERVER_PASSWORD"),
			envy.Get("SQL_SERVER_PORT"),
			envy.Get("SQL_SERVER_DB_NAME_1"),
		),
	},
}
