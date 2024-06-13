package db

import (
	"cms/utils/env"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func getDBConnectionString() string {
	host := env.Get("DB_HOST", "localhost")
	port := env.Get("DB_PORt", "5432")
	user := env.Get("DB_USER", "USER")
	password := env.Get("DB_PASSWORD", "password")
	dbname := env.Get("DB_NAME", "dbname")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
}

func ConnectDB(maxOpenConn, maxIdleConn, maxLifetime int) (*sql.DB, error) {
	connStr := getDBConnectionString()

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConn)
	db.SetMaxIdleConns(maxIdleConn)
	db.SetConnMaxLifetime(time.Duration(maxLifetime) * time.Minute)

	return db, nil
}
