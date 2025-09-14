package main

import (
	"log"
	"net/http"
	"url-shortener/config"
	"url-shortener/handlers"
	"url-shortener/repository"
	"url-shortener/service"

	"github.com/redis/go-redis/v9"
)

func main() {
	db := config.InitDB() // returns *sql.DB
	repo := repository.NewURLRepository(db)

	redisClient := redis.NewClient(&redis.Options{
		Addr: "redis:6379", // docker-compose hostname
	})

	urlService := service.NewURLService(repo, redisClient)
	handler := handlers.NewHandler(urlService)

	log.Println("Starting server on :8080")
	r := handler.Routes()
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
