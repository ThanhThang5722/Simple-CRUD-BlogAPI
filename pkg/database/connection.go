package database

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var instance *sql.DB
var err error

func ConnectMySQL() (*sql.DB, error) {
	dsn := os.Getenv("DB_ConnectLink")
	instance, err = sql.Open("mysql", dsn)
	return instance, err
}
func GetInstance() *sql.DB {
	return instance
}
