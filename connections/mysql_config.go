package connections

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

var Db *sql.DB

func InitMysql() (err error) {
	dsn := "root:147526@tcp(127.0.0.1:3306)/tree?charset=utf8mb4&parseTime=True"
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = Db.Ping()
	if err != nil {
		return err
	}
	log.Info("MySQL Init Success！")
	return nil
}
