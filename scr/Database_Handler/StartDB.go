package databasehandler

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

func StartDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error Opening DB: %v\n", err)
		return db, err
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(60 * time.Minute)

	if err := db.Ping(); err != nil {
		log.Fatalf("DB ping failed:v %v \n", err)
		return db, err
	}
	fmt.Println("Sucessfully connectedto MySQL!")

	return db, nil
}
