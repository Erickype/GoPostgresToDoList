package main

import (
	"fmt"
	"github.com/Erickype/GoPostgresToDoList/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//Routes
	app.Get("/", routes.GetHandler)

	app.Post("/", routes.PostHandler)

	app.Put("/update", routes.UpdateHandler)

	app.Delete("/delete", routes.DeleteHandler)

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}
