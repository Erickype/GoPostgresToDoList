package routes

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
)

func GetHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Get")
}
func PostHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Post")
}
func UpdateHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Update")
}
func DeleteHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Delete")
}
