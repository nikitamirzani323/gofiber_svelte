package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nikitamirzani323/gofiber_svelte/configs"
)

var db *sql.DB
var err error

func InitMysql() {
	conf := configs.GetConfMysql()
	connectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME + "?parseTime=true"

	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Panic("koneksi string error")
	}

	err = db.Ping()
	if err != nil {
		log.Panic("DNS invalid")
	}
}

func CreateCon() *sql.DB {
	return db
}
