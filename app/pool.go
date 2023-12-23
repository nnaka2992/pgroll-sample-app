package main

import (
	"context"
	"fmt"
	"os"
	"log"
	"strconv"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)
var pool *pgxpool.Pool

func NewPool(config *pgxpool.Config) (*pgxpool.Pool, error) {
	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Println("Error connecting to database: ", err)
		return nil, err
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, err
	}

	return pool, nil
}

func NewPoolFromEnv() (*pgxpool.Pool, error) {
	user := os.Getenv("DB_USER")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	db := os.Getenv("DB_NAME")
	maxConn, err := strconv.ParseInt(os.Getenv("DB_MAX_CONN"),10,32)
	if err != nil {
		log.Println("Error parsing max connection: ", err)
		return nil, err
	}
	searchPath := os.Getenv("DB_SEARCH_PATH")
	// format string
	connString := fmt.Sprintf("postgres://%s@%s:%s/%s?sslmode=disable", user, host, port, db)

	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		log.Println("Error parsing connection string: ", err)
		return nil, err
	}
	config.MaxConns = int32(maxConn)
	config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		if _, err := conn.Exec(ctx, fmt.Sprintf("SET search_path TO %s", searchPath)); err != nil{
			return err
		}
		return nil
	}

	return NewPool(config)
}

func NewPoolFromURL() (*pgxpool.Pool, error) {
	pool, err := pgxpool.Connect(context.Background(), "postgres://postgres:postgres@localhost:5432/crudapp")
	if err != nil {
		log.Println("Error connecting to database: ", err)
		return nil, err
	}
	return pool, nil
}
