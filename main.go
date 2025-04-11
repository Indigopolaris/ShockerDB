package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	// Initialize a new Fiber app
	app := fiber.New()

	// Define a route for the GET method on the root path '/'
	app.Get("/", func(c fiber.Ctx) error {
		// Send a string response to the client
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/Login", func(c fiber.Ctx) error {
		//login page
		msg := fmt.Sprintf("Login Endpoint")
		return c.SendString(msg)
	})

	app.Get("/Dashboard", func(c fiber.Ctx) error {
		//dashboard for the power readings, this will have the power readings over the month for each month
		msg := fmt.Sprintf("Dashboard Endpoint")
		return c.SendString(msg)
	})

	app.Get("/Form", func(c fiber.Ctx) error {
		//Input form for power readings
		msg := fmt.Sprintf("Input Data Endpoint")
		return c.SendString(msg)
	})

	app.Get("/DataViewer", func(c fiber.Ctx) error {
		//view all historical data
		msg := fmt.Sprintf("Explore Data Endpoint")
		return c.SendString(msg)
	})

	app.Get("/api/*", func(c fiber.Ctx) error {
		msg := fmt.Sprintf("API Endpoint", c.Params("*"))
		return c.SendString(msg)
	})

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
