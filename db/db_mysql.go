package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nikitamirzani323/gofiber_svelte/configs"
)

var db *sql.DB
var err error

func InitMysql() {
	conf := configs.ConfMysql()
	connectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME

}
