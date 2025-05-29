package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("failed to open db:", err)
	}
	defer db.Close()

	err = db.PingContext(context.Background())
	if err != nil {
		log.Fatal("failed to ping db:", err)
	}

	fmt.Println("Connected to DB successfully")

	// Example query to check the table exists
	rows, err := db.QueryContext(context.Background(), "SELECT id, name FROM users")
	if err != nil {
		log.Fatal("query failed:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal("scan failed:", err)
		}
		fmt.Printf("User: %d %s\n", id, name)
	}
}
