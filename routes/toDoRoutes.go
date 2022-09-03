package routes

import (
	"github.com/gofiber/fiber/v2"
)

func GetHandler(c *fiber.Ctx) error {
	return c.SendString("Get")
}
func PostHandler(c *fiber.Ctx) error {
	return c.SendString("Post")
}
func UpdateHandler(c *fiber.Ctx) error {
	return c.SendString("Update")
}
func DeleteHandler(c *fiber.Ctx) error {
	return c.SendString("Delete")
}
