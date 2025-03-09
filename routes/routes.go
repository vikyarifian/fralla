package routes

import (
	"fralla/auth"
	"fralla/handlers"
	"fralla/templ/pages"
	"fralla/utils"

	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App) {
	DashboardRoute(app)
	PageRoute(app)
	HandleRoute(app)

	app.Get("/logout", func(c *fiber.Ctx) error {
		auth.ClearSession(c)
		c.Response().Header.Set("HX-Redirect", "/")
		return c.Redirect("/")
	})

	app.Use(func(c *fiber.Ctx) error {
		return utils.Render(c, pages.NotFound(auth.IsAuthenticated(c)))
	})

}

func DashboardRoute(app *fiber.App) {

}

func PageRoute(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return utils.Render(c, pages.FrontPage2(auth.IsAuthenticated(c)))
	})
	app.Get("/401", func(c *fiber.Ctx) error {
		return utils.Render(c, pages.Forbidden(auth.IsAuthenticated(c)))
	})
	app.Get("/about", func(c *fiber.Ctx) error {
		return utils.Render(c, pages.About(auth.IsAuthenticated(c)))
	})
	app.Get("/shop", func(c *fiber.Ctx) error {
		return utils.Render(c, pages.Shop(auth.IsAuthenticated(c)))
	})
}

func HandleRoute(app *fiber.App) {
	app.Post("/register", handlers.HandleRegister)
	app.Post("/login", handlers.HandleLogin)
	app.Get("/hide-popup", handlers.HidePopUp)
}
