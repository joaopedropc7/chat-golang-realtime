package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pusher/pusher-http-go/v5"
)

func main() {
	app := fiber.New()

	app.Use(cors.New()) // default config

	pusherClient := pusher.Client{
		AppID:   "1696035",
		Key:     "39b59267b302534cb4f8",
		Secret:  "f89ff53fc902742a0760",
		Cluster: "sa1",
		Secure:  true,
	}

	app.Post("/api/messages", func(c *fiber.Ctx) error {
		var data map[string]string

		if err := c.BodyParser(&data); err != nil {
			return err
		}

		err := pusherClient.Trigger("chat", "message", data)
		if err != nil {
			fmt.Println(err.Error())
		}

		return c.JSON([]string{})
	})

	app.Listen(":8000")
}
