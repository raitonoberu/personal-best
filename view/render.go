package view

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v3"
)

func Render(c fiber.Ctx, template templ.Component) error {
	c.Set("Content-type", "text/html")
	return template.Render(c.Context(), c.Response().BodyWriter())
}
