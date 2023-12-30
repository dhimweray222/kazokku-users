package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"
)

func NewPostgresPool() *pgxpool.Pool {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	minConns := os.Getenv("DB_POOL_MIN")
	maxConns := os.Getenv("DB_POOL_MAX")

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", user, password, host, port, dbName)

	poolConfig, err := pgxpool.ParseConfig(dsn)

	if err != nil {
		log.Fatal(err)
	}

	minConnsInt, err := strconv.Atoi(minConns)
	if err != nil {
		log.Fatalf("DB_MIN_CONNS expected to be integer, got %s", minConns)
	}
	maxConnsInt, err := strconv.Atoi(maxConns)
	if err != nil {
		log.Fatalf("DB_MAX_CONNS expected to be integer, got %s", maxConns)
	}

	poolConfig.MinConns = int32(minConnsInt)
	poolConfig.MaxConns = int32(maxConnsInt)

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	log.Println(pool)
	if err != nil {
		log.Fatal(err)
	}

	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := pool.Ping(c); err != nil {
		log.Fatal(err)
	}

	log.Println("Database connected", dsn)
	return pool
}
