package database

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *sql.DB
var GDB *gorm.DB

func GetConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true",
		os.Getenv("DB_USER"), os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))
}

func InitDB() error {
	conn, err := sql.Open("mysql", GetConnectionString())
	if err != nil {
		return err
	}
	DB = conn

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: conn,
	}))
	if err != nil {
		return err
	}

	GDB = db

	return nil
}
