package routes

import (
	"database/sql"
	"github.com/Erickype/GoPostgresToDoList/models"
	"github.com/gofiber/fiber/v2"
	"log"
)

func GetHandler(c *fiber.Ctx, db *sql.DB) error {
	var res string
	var todos []string

	q := `SELECT * FROM todos`

	rows, err := db.Query(q)

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		err := rows.Scan(&res)
		if err != nil {
			return err
		}
		todos = append(todos, res)
	}

	return c.Render("index", fiber.Map{"Todos": todos})
}
func PostHandler(c *fiber.Ctx, db *sql.DB) error {

	newTodo := models.Todo{}

	err := c.BodyParser(&newTodo)
	if err != nil {
		log.Printf("An error ocurred: %d", err)
		return c.SendString(err.Error())
	}

	return c.Redirect("/")
}
func UpdateHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Update")
}
func DeleteHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Delete")
}
