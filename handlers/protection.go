package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nurtidev/network-protection-assessment/models"
)

func GetProtectionMethods(c *fiber.Ctx) error {
	return c.JSON(models.ProtectionMethods)
}

func SimulateProtection(c *fiber.Ctx) error {
	// Логика для проведения симуляции атак
	result := map[string]float64{
		"Антивирус": 0.85,
		"Фаервол":   0.90,
	}
	return c.JSON(result)
}
