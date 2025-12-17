package main

import (
	"context"
	"fmt"
	db "go-backend-task/db/sqlc"
	"go-backend-task/internal/handler"
	"go-backend-task/internal/service"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

// Update this connection string with your actual DB credentials
const dbSource = "postgresql://postgres:root@localhost:5432/users_db?sslmode=disable"

func main() {
	// 1. Initialize Zap Logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// 2. Connect to Database using pgx
	conn, err := pgx.Connect(context.Background(), dbSource)
	if err != nil {
		logger.Fatal("Unable to connect to database", zap.Error(err))
	}
	defer conn.Close(context.Background())

	// 3. Initialize Architecture Layers
	queries := db.New(conn)
	userService := service.NewUserService(queries)
	userHandler := handler.NewUserHandler(userService, logger)

	// 4. Setup Fiber App
	app := fiber.New()

	// 5. Register Routes
	api := app.Group("/users")
	
	api.Post("/", userHandler.CreateUser)
	api.Get("/", userHandler.ListUsers)
	api.Get("/:id", userHandler.GetUser)
	api.Put("/:id", userHandler.UpdateUser)
	api.Delete("/:id", userHandler.DeleteUser)

	// 6. Start Server
	fmt.Println("Server running on port 3000")
	log.Fatal(app.Listen(":3000"))
}