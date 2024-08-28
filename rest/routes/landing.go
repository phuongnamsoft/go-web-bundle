package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/phuongnamsoft/go-web-bundle/app"
	"github.com/phuongnamsoft/go-web-bundle/rest/controllers"
	"github.com/phuongnamsoft/go-web-bundle/rest/middlewares"
)

func LandingRoutes(web fiber.Router) {
	web.Get("/", controllers.Landing)
	web.Get("/ping", Pong)
	web.Get("/all-routes", AllRoutes)
	web.Get("/do/verify-email", middlewares.ValidateConfirmToken, controllers.VerifyRegisteredEmail)
}

func Pong(c *fiber.Ctx) error {
	return c.SendString("Pong")
}

func AllRoutes(c *fiber.Ctx) error {
	return c.JSON(app.Http.Server.Stack())
}
