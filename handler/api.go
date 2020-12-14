package handler

import "github.com/gofiber/fiber/v2"

func Index(ctx *fiber.Ctx) error {
	return ctx.Render("index", fiber.Map{"index": true})
}

func Whisper(ctx *fiber.Ctx) error {
	return ctx.Render("whisper", fiber.Map{"whisper": true})
}

func Leacots(ctx *fiber.Ctx) error {
	return ctx.Render("leacots", fiber.Map{"leacots": true})
}

func Album(ctx *fiber.Ctx) error {
	return ctx.Render("album", fiber.Map{"album": true})
}

func User(ctx *fiber.Ctx) error {
	return ctx.Render("user", fiber.Map{"user": true})
}

func About(ctx *fiber.Ctx) error {
	return ctx.Render("about", fiber.Map{"about": true})
}
