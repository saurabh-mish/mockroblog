package db

import (
	"testing"
	"mockroblog/pkg/db"
	"golang.org/x/exp/slices"
)


func TestGetRepository(t *testing.T) {
	sqlitedb := db.GetRepository()

	t.Run("ping database", func(t *testing.T) {
		got := sqlitedb.Ping()
		if got != nil {
			t.Error("Unable to ping sqlite database ...")
		}
	})

	t.Run("check if tables exist", func(t *testing.T) {
		rows, err := sqlitedb.Query("SELECT name FROM sqlite_master WHERE type='table'")
		if err != nil {
			t.Errorf("Error performing query to retrieve tables: %v", err)
		}

		defer rows.Close()

		var table string
		var tables []string

		for rows.Next() {
			err := rows.Scan(&table)
			if err != nil {
				t.Error(err)
			}
			tables = append(tables, table)
		}

		if !slices.Contains(tables, "user") || !slices.Contains(tables, "post") {
			t.Errorf("Tables 'user' and 'post' are not present in the database")
		}
	})
}
