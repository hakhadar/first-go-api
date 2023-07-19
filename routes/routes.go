package routes

import (
	"gobackReactfront/go_auth/conttollers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", conttollers.Register)
	app.Post("/api/login", conttollers.Login)
	app.Get("/api/user", conttollers.User)
	app.Post("/api/logout", conttollers.Logout)

}
