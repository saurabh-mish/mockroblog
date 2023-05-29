package db

import (
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
)

const (
	driver = "sqlite"
	source = "../../init/mockroblog.db"
)

func GetRepository() *sql.DB {
	db, err := sql.Open(driver, source)
	if err != nil {
		fmt.Printf("Unable to open source database: %v", err)
	}
	return db
}
