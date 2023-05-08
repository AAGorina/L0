package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

var Conn *pgxpool.Pool

func DBConnection() {
	var err error
	Conn, err = pgxpool.New(context.Background(),
		"postgres://monokuma:monokuma@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		fmt.Println("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
}
