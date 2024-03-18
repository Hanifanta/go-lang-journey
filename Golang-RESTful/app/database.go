package app

import (
	"database/sql"
	"golang-restful/helper"
	"time"
)

func NewDB() *sql.DB {
	//db, err := sql.Open("mysql", "root:amikompedia@tcp(localhost:3306)/golang_restful")
	db, err := sql.Open("mysql", "root:amikompedia@tcp(localhost:3306)/golang_migration")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
