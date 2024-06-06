package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nurtidev/network-protection-assessment/charts"
	"github.com/nurtidev/network-protection-assessment/handlers"
	"log"
)

func main() {
	app := fiber.New()

	app.Get("/methods", handlers.GetProtectionMethods)
	app.Post("/simulate", handlers.SimulateProtection)
	app.Get("/chart", charts.GenerateChart)

	log.Fatal(app.Listen(":3000"))
}
