package router

import (
	"github.com/RianIhsan/goCloudinary/controllers"
	"github.com/gofiber/fiber/v2"
)

func InitRoute(app *fiber.App) {
	app.Post("/user", controllers.User)
}
