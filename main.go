package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/francocastro-94/Crud-Fiber/db"
	"github.com/francocastro-94/Crud-Fiber/model"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {

	// Initialize standard Go html template engine
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	//Show index template with all products
	app.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", fiber.Map{
			"Products": db.GetAll(),
		})
	})

	app.Get("/form", func(c *fiber.Ctx) error {
		// Render form template
		return c.Render("formCreate", nil)
	})

	//create product
	app.Post("/createNew", func(c *fiber.Ctx) error {
		newProduct := model.Product{
			Name:        c.FormValue("name"),
			Image:       c.FormValue("image"),
			Description: c.FormValue("description"),
		}
		db.SaveOneProduct(&newProduct)
		return c.Redirect("/")
	})

	//Update a product (use POST method by using mvc  )
	app.Post("/update", func(c *fiber.Ctx) error {
		newProduct := model.Product{
			Name:        c.FormValue("name"),
			Image:       c.FormValue("image"),
			Description: c.FormValue("description"),
		}
		db.UpdateOneById(c.FormValue("id"), &newProduct)
		return c.Redirect("/")
	})

	app.Get("/updateForm/:id", func(c *fiber.Ctx) error {
		return c.Render("updateForm", fiber.Map{
			"ID": c.Params("id"),
		})
	})

	//Delete a product (use GET method by using mvc  )
	app.Get("/delete/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		db.DeleteById(id)
		return c.Redirect("/")
	})
	//Show a product
	app.Get("/show/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		product := db.GetOneById(id)
		fmt.Println(product)
		return c.Render("show", fiber.Map{
			"Product": product,
		})
	})

	log.Fatal(app.Listen(":8080"))
}
