package main

import (
	"context"
	_ "github.com/jackc/pgx/v5/pgtype"
	userv1 "github.com/kirigaikabuto/bike-main-api/gen/proto"
	grpc2 "github.com/kirigaikabuto/bike-main-api/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
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

	grpcServer := grpc.NewServer()
	userServer := &grpc2.UserServer{Queries: queries}

	userv1.RegisterUserServiceServer(grpcServer, userServer)

	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("🚀 gRPC server running on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("gRPC server failed: %v", err)
	}
}
