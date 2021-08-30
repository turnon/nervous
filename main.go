package main

import (
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/turnon/nervous/views"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		bytes, _ := views.Render("cal.html", nil)
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		return c.Send(bytes)
	})

	app.Get("/lib/:filename", func(c *fiber.Ctx) error {
		filename := c.Params("filename")
		if filepath.Ext(filename) == ".css" {
			c.Set(fiber.HeaderContentType, "text/css")
		} else {
			c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJavaScript)
		}
		bytes := views.Lib(filename)
		return c.Send(bytes)
	})

	app.Listen(":3000")
}
