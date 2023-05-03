package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lnzx/ntfy/internal/sms"
	"log"
	"os"
)

var API_KEY = "Test1234##"

func init() {
	apiKey := os.Getenv("API_KEY")
	if apiKey != "" {
		API_KEY = apiKey
	}
}

func main() {
	app := fiber.New()

	app.Post("/sms", func(c *fiber.Ctx) error {
		if c.Get("Authorization") != API_KEY {
			return fiber.ErrUnauthorized
		}
		body := new(Body)
		if err := c.BodyParser(body); err != nil {
			return fiber.ErrBadRequest
		}
		if body.Mobile == "" || body.Event == "" {
			return fiber.ErrBadRequest
		}
		if err := sms.Send(body.Mobile, body.Event); err != nil {
			return err
		}
		return nil
	})

	log.Fatal(app.Listen(":8080"))
}

type Body struct {
	Mobile string
	Event  string
}
