package conf

import "github.com/gofiber/fiber/v2"

func InitFiber() *fiber.App {
	c := fiber.Config{}
	c.Prefork = preforkFlag
	app := fiber.New(c)
	return app
}
