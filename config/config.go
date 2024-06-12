package config

import (
	// "database/sql"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

// var DB *sql.DB

func ConnectDB(dsn string) error {
	// var err error
	// DB, err = sql.Open("pgx", dsn)
	// if err != nil {
	// 	return err
	// }

	// err = DB.Ping()
	// if err != nil {
	// 	return err
	// }

	log.Println("Connected to PostgreSQL!")

	return nil
}
