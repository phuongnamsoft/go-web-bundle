package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/phuongnamsoft/go-web-bundle/rest/middlewares"
)

func LoadRoutes(app *fiber.App) {
	api := app.Group("api").Use(middlewares.AuthApi())
	web := app.Group("")
	ApiRoutes(api)
	WebRoutes(web)
}
