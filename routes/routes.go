package routes

import (
	"fralla/auth"
	"fralla/templ/pages"
	"fralla/utils"

	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App) {
	DashboardRoute(app)
	PageRoute(app)

	app.Get("/logout", func(c *fiber.Ctx) error {
		auth.ClearSession(c)
		c.Response().Header.Set("HX-Redirect", "/")
		return c.Redirect("/")
	})
}

func DashboardRoute(app *fiber.App) {

}

func PageRoute(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return utils.Render(c, pages.FrontPage(auth.IsAuthenticated(c)))
	})
}
