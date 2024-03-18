package Golang_Database_MySQL

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestEmpty(t *testing.T) {
	// your test
}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:amikompedia@tcp(localhost:3306)/belajar_golang_database")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
