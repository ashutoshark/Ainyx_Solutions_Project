package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/user-api/config"
	"github.com/user-api/db/sqlc"
	"github.com/user-api/internal/handler"
	"github.com/user-api/internal/repository"
	"github.com/user-api/internal/routes"
	"github.com/user-api/internal/service"
)

func main() {
	// load config
	cfg := config.Load()

	// database connection
	db, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}
	defer db.Close()
	log.Println("connected to db!")

	// setup everything
	q := sqlc.New(db)
	repo := repository.NewUserRepository(q)
	srv := service.NewUserService(repo)
	h := handler.NewUserHandler(srv)

	// create app
	app := fiber.New()
	routes.Setup(app, h)

	// start
	log.Println("server starting on port", cfg.Port)
	log.Fatal(app.Listen(":" + cfg.Port))
}
