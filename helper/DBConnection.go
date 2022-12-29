package helper

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/kehadiran?parseTime=true")
	CatchPanic(err)

	db.SetMaxIdleConns(3)
	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(10 * time.Minute)

	return db
}
