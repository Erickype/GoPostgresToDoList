package main

import (
	"fmt"
	"github.com/Erickype/GoPostgresToDoList/database"
	"github.com/Erickype/GoPostgresToDoList/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//Database
	db := database.GetConnection()

	//Routes
	app.Get("/", func(ctx *fiber.Ctx) error {
		return routes.GetHandler(ctx, db)
	})

	app.Post("/", func(ctx *fiber.Ctx) error {
		return routes.PostHandler(ctx, db)
	})

	app.Put("/update", func(ctx *fiber.Ctx) error {
		return routes.UpdateHandler(ctx, db)
	})

	app.Delete("/delete", func(ctx *fiber.Ctx) error {
		return routes.DeleteHandler(ctx, db)
	})

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}
