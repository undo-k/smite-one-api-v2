package tools

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
)

var (
	host     = os.Getenv("DB_HOST")
	port, _  = strconv.Atoi(os.Getenv("DB_PORT"))
	user     = os.Getenv("PG_DB_USER")
	password = os.Getenv("PG_DB_PW")
	dbname   = os.Getenv("DB_NAME")
)

func OpenConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
