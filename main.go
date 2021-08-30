package main

import (
	"encoding/json"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/turnon/nervous/views"
)

type jsonObject struct {
	obj interface{}
}

func (j *jsonObject) EvalObject() string {
	bytes, _ := json.Marshal(j.obj)
	return "<pre style='display:none'>" +
		string(bytes) +
		"</pre><script>var obj = JSON.parse(document.querySelector('pre').innerText)</script>"
}

type calendarPage struct {
	Events []event `json:"events"`
}

type event struct {
	Title string `json:"title"`
	Start string `json:"start"`
	End   string `json:"end"`
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		calPage := calendarPage{Events: []event{
			{Title: "aaa", Start: "2021-08-21"},
			{Title: "aaa", Start: "2021-08-25", End: "2021-08-27"},
		}}

		bytes, _ := views.Render("cal.html", &jsonObject{calPage})
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
