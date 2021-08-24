package driver

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	dsn := os.Getenv("DATABASE_URL")

	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println("Database connection failed")
	}

	err = conn.Ping()
	if err != nil {
		log.Fatal("DB connection not established => ", err)
	}

	log.Println("Database connected successfully")

	return conn
}
