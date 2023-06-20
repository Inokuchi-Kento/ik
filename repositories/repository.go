package repositories

import (
	"context"
	"fmt"
	"log"

	"blrpc/config"
	"blrpc/ent"

	_ "github.com/lib/pq"
)

var (
	client *ent.Client
	err    error
)

func OpenDB(ctx context.Context, cfg *config.Config) *ent.Client {
	log.Printf("DB_config: %v", cfg)
	client, err = ent.Open("postgres", fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort,
	))
	if err != nil {
		log.Fatalf("failed connecting to Postgres: %v", err)
	}

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}
