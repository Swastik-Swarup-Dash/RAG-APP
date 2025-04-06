package db

import (
	"context"
	"os"
	
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pgvector/pgvector-go"
)

// Declare exported DB variable
var DB *pgxpool.Pool

func InitDB() error {
	connStr := os.Getenv("DB_CONN")
	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return err
	}
	DB = pool
	return nil
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}

func StoreEmbedding(content string, embedding []float32) error {
	_, err := DB.Exec(context.Background(),
		"INSERT INTO documents (content, embedding) VALUES ($1, $2)",
		content,
		pgvector.NewVector(embedding),
	)
	return err
}
