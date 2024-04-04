package main

import(
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
}

func main() {
	err := godotenv.Load()

	if err!= nil {
		fmt.Println(err)
	}

	app := fiber.New

	app.Use(logger.New())

	setupRoutes(app)
}