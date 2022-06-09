package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	HOST = "localhost"
	PORT = "5432"
)

var ErrNoMatch = fmt.Errorf("No matching record")

type Database struct {
	Conn *sql.DB
}

func Initialize() (Database, error) {
	db := Database{}
	daba := fmt.Sprintf("host =%s port =%s user =%s password =%s dbname =%s sslmode=disable",
		HOST, PORT, username, password, database)
	conn, err := sql.Open("postgres", daba)
	if err != nil {
		return db, nil
	}
	db.Conn = conn
	err = db.Conn.Ping()
	if err != nil {
		return db, err
	}
	log.Println("Database connected")
	return db, nil
}
