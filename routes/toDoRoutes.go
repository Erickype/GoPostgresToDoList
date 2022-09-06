package routes

import (
	"database/sql"
	"fmt"
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

	var q = `INSERT INTO todos VALUES ($1)`
	newTodo := models.Todo{}

	err := c.BodyParser(&newTodo)
	if err != nil {
		log.Printf("An error ocurred: %d", err)
		return c.SendString(err.Error())
	}

	fmt.Printf("New item: %v\n", newTodo)
	//Insert to database
	if newTodo.Item != "" {
		_, err := db.Exec(q, newTodo.Item)
		if err != nil {
			log.Fatalf("An error ocurred while executing query: %v\n", err)
		}
	}

	return c.Redirect("/")
}

func UpdateHandler(c *fiber.Ctx, db *sql.DB) error {

	oldItem := c.Query("oldItem")
	newItem := c.Query("newItem")
	q := `UPDATE todos SET item=$1 WHERE item=$2`

	exec, err := db.Exec(q, newItem, oldItem)
	if err != nil {
		log.Fatalf("An error ocurred while executing query: %v\n%v", err, exec)
	}

	return c.SendString("Updated")
}

func DeleteHandler(c *fiber.Ctx, db *sql.DB) error {

	item := c.Query("item")
	q := `DELETE FROM todos WHERE item = $1`

	exec, err := db.Exec(q, item)
	if err != nil {
		log.Fatalf("An error ocurred while executing query: %v\n%v", err, exec)
	}

	return c.SendString("Delete")
}
