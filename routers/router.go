package routers

import (
	"free-blog/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetRouters(app *fiber.App) {
	app.Get("/", handler.Index)
	app.Get("/whisper", handler.Whisper)
	app.Get("/leacots", handler.Leacots)
	app.Get("/album", handler.Album)
	app.Get("/about", handler.About)
	app.Get("/user", handler.User)

	// api
	api:=app.Group("api",logger.New())
	api.Post("/login",handler.Login)
}
