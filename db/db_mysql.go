package db

import (
	"database/sql"

	"bitbucket.org/isbtotogroup/api_go/config"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Init() {
	conf := config.GetConfigMysql()
	connectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME + "?net_write_timeout=6000"

	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		panic("koneksi string serror")
	}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(4)
	// db.SetConnMaxLifetime(time.Minute * 4)
	err = db.Ping()
	if err != nil {
		panic("DNS invalid")
	}
}
func CreateCon() *sql.DB {
	return db
}
