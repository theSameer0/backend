package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	dsn := "user=sameermishra password=sameer dbname=mydatabase port=5430 sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	CheckErr(err)
	return db
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
