package connector

import (
	"database/sql"
	"database/sql/driver"

	"github.com/go-sql-driver/mysql"
)

func getConnector() (driver.Connector, error) {
	config := BuildConfig()
	connector, err := mysql.NewConnector(config)
	return connector, err
}

func GetDb() (*sql.DB, error) {
	var db *sql.DB
	var err error
	if connector, err := getConnector(); err == nil {
		db = sql.OpenDB(connector)
	}
	return db, err
}
