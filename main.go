package main

import (
	"encoding/json"
	"github.com/arsmn/fiber-swagger"
	"github.com/gofiber/fiber"
	"github.com/idawud/bookstore/book"
	"github.com/idawud/bookstore/database"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"log"
)

func setupRoutes(app *fiber.App)  {
	app.Get("/api/v1/book", book.GetAllBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.AddBook)
	app.Put("/api/v1/book/:id", book.UpdateBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
	app.Get("/docs", showDocs)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("Failed To Open Connection: " + err.Error())
	}
	log.Println("DB connection Successful")

	database.DBConn.AutoMigrate(&book.Book{})
	log.Println("Migration Completed")
}


// @title BookStore API
// @version 1.0
// @description This is a swagger for my bookstore api
// @termsOfService http://swagger.io/terms/
// @contact.name idawud
// @contact.email ismaildawud96@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func main() {
	app := fiber.New()

	initDatabase()
	defer database.DBConn.Close()

	setupRoutes(app)

	app.Use("/swagger", swagger.New(swagger.Config{ // custom
		URL: "http://localhost:3000/docs",
		DeepLinking: false,
	}))
	app.Listen(3000)
}

func showDocs(c *fiber.Ctx) {
	raw, err := ioutil.ReadFile("./docs/swagger.json")
	var result map[string]interface{}
	json.Unmarshal([]byte(raw), &result)
	if err != nil{
		c.Send("Erroe")
	}
	c.JSON(result)
}

