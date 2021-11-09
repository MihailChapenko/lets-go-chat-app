package db

import (
	"context"
	"fmt"
	"github.com/MihailChapenko/lets-go-chat-app/internal/config"
	"github.com/jackc/pgx/v4"
	"os"
)

var conn *pgx.Conn

func Init(cfg *config.DB) {
	conn, err := pgx.Connect(context.Background(), cfg.DataSource)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
}

// GetDB get database connection
func GetDB() *pgx.Conn {
	return conn
}
