package database

import (
	"database/sql"

	// needed by database/sql
	_ "github.com/go-sql-driver/mysql"
)

// NewDB func
func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:navri@/goblog?parseTime=true")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	// db.SetConnMaxLifetime(time.Minute * 3)
	// db.SetMaxOpenConns(10)
	// db.SetMaxIdleConns(10)

	return db
}
