package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pgvector/pgvector-go"
)

var DB *pgxpool.Pool // Ensure this is exported

func InitDB() error {
	connStr := os.Getenv("DB_CONN")
	if connStr == "" {
		return fmt.Errorf("DB_CONN environment variable not set")
	}

	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return fmt.Errorf("failed to parse DB config: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return fmt.Errorf("failed to create connection pool: %w", err)
	}

	// Test the connection
	if err := pool.Ping(context.Background()); err != nil {
		return fmt.Errorf("database ping failed: %w", err)
	}

	DB = pool
	return nil
}

func StoreEmbedding(content string, embedding []float32) error {
	if DB == nil {
		return fmt.Errorf("database connection not initialized")
	}

	_, err := DB.Exec(context.Background(),
		"INSERT INTO documents (content, embedding) VALUES ($1, $2)",
		content,
		pgvector.NewVector(embedding),
	)
	return err
}
