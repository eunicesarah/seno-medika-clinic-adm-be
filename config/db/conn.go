package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"seno-medika.com/config/variable"
)

var DB = Conn()

func Conn() *sql.DB {
	postgresInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		variable.DbHost, 5432, variable.DbUser, variable.DBPass, variable.DBName)

	db, err := sql.Open("postgres", postgresInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
