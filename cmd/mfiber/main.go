package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New(fiber.Config{
		Prefork: true,
	})
	app.Get("/found", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/plain")
		return c.Status(200).SendString("Route /found")
	})
	println("Listening on :8080")
	// panic(app.Server().ListenAndServe(":8080")) // Why returns "Cannot GET /found"?
	panic(app.Listen(":8080"))
}
