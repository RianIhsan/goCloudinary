package main

import (
	"github.com/RianIhsan/goCloudinary/database"
	"github.com/RianIhsan/goCloudinary/router"
	"github.com/gofiber/fiber/v2"
)

func main() {

	database.SetupDatabase()
	database.RunMigrate()

	app := fiber.New()

	router.InitRoute(app)
	app.Listen(":8080")
}
