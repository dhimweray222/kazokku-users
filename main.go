package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dhimweray222/users/config"
	"github.com/dhimweray222/users/controller"
	"github.com/dhimweray222/users/repository"
	"github.com/dhimweray222/users/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	time.Local = time.UTC
	db := config.NewPostgresPool()

	serverHost := os.Getenv("SERVER_URI")
	serverPort := os.Getenv("SERVER_PORT")

	userStore := repository.NewUserStore(db)
	userRepository := repository.NewUserRepository(userStore)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	app := fiber.New(fiber.Config{BodyLimit: 10 * 1024 * 1024})
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "*",
		AllowHeaders:     "*",
		AllowCredentials: true,
	}))

	userController.NewUserRouter(app)

	host := fmt.Sprintf("%s:%s", serverHost, serverPort)
	err := app.Listen(host)

	log.Println(err)
}
