package main

import (
	"encoding/json"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/turnon/nervous/db"
	"github.com/turnon/nervous/ent"
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
	Events []*db.Event `json:"events"`
}

type newEvents struct {
	Dates      []string `form:"dates[]"`
	Tag        string   `form:"tag"`
	Name       string   `form:"name"`
	Continuous int      `form:"continue"`
}

func main() {
	app := fiber.New()
	var dbHandler db.DbHandler = &ent.DbHandler{}

	app.Get("/", func(c *fiber.Ctx) error {
		calPage := calendarPage{}

		start_at := c.Query("start", time.Now().Format("2006-01-02"))
		hv := strings.Split(c.Query("layout", "1x1"), "x")
		h, _ := strconv.Atoi(hv[0])
		v, _ := strconv.Atoi(hv[1])
		events := dbHandler.LoadEvents(start_at, h*v)
		calPage.Events = events

		bytes, _ := views.Render("cal.html", &jsonObject{calPage})
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		return c.Send(bytes)
	})

	app.Post("/new", func(c *fiber.Ctx) error {
		ne := newEvents{}
		c.BodyParser(&ne)
		dbHandler.NewEvents(ne.Dates, ne.Tag, ne.Name, ne.Continuous)
		return c.Redirect("/")
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
