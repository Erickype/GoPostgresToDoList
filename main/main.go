package main

import (
	"fmt"
	"github.com/Erickype/GoPostgresToDoList/rutas"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//Routes
	app.Get("/", rutas.GetHandler)

	app.Post("/", rutas.PostHandler)

	app.Put("/update", rutas.UpdateHandler)

	app.Delete("/delete", rutas.DeleteHandler)

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}
