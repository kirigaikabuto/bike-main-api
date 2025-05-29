package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"

	"github.com/kirigaikabuto/bike-main-api/internal/db" // replace with your actual module name
)

func main() {
	_ = godotenv.Load()
	ctx := context.Background()

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		log.Fatalf("Unable to create pgx pool: %v", err)
	}
	defer pool.Close()

	conn, err := pool.Acquire(ctx)
	if err != nil {
		log.Fatalf("Unable to acquire connection from pool: %v", err)
	}
	defer conn.Release()

	queries := db.New(conn)

	user, err := queries.CreateUser(ctx, db.CreateUserParams{
		Name:  "Yerassyl",
		Email: pgtype.Text{},
	})
	if err != nil {
		log.Fatalf("Error creating user: %v", err)
	}

	fmt.Printf("Created user: %+v\n", user)

	fetchedUser, err := queries.GetUserByID(ctx, user.ID)
	if err != nil {
		log.Fatalf("Error fetching user: %v", err)
	}
	fmt.Printf("Fetched user: %+v\n", fetchedUser)

	_, err = queries.CreateBook(ctx, db.CreateBookParams{
		Name:  "book1",
		Price: 3.4,
	})
	if err != nil {
		log.Fatalf("Error creating book: %v", err)
	}

	books, err := queries.ListBooks(ctx)
	if err != nil {
		log.Fatalf("Error listing books: %v", err)
	}
	fmt.Printf("Listed books: %+v\n", books)
}
