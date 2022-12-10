package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbUser    = "go"
	dbPass    = "dev@Golang2022"
	dbName    = "devbook"
	urlParams = "charset=utf8&parseTime=True&loc=Local"
)

func connection() (*sql.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@/%s?%s", dbUser, dbPass, dbName, urlParams)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}

func Connect() (*sql.DB, error) {
	return connection()
}
