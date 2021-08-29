package main

import (
	"embed"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

//go:embed static
var static embed.FS

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		bytes := readStatic("cal.html")
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
		bytes := readStatic("lib/" + filename)
		return c.Send(bytes)
	})

	app.Listen(":3000")
}

func readStatic(name string) []byte {
	bytes, err := static.ReadFile("static/" + name)
	if err != nil {
		panic(err)
	}
	return bytes
}
