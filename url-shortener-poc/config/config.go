package config

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./data/urls.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
	return rdb
}
