package book

import (
	"github.com/gofiber/fiber"
	"github.com/idawud/bookstore/database"
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Title string `json:"title"`
	Author string `json:"author"`
	Rating int `json:"rating"`
}

type HTTPError struct {
	status  string
	message string
}

// GetAllBooks godoc
// @Summary Show all books
// @Description list out all books available
// @Accept  json
// @Produce  json
// @Success 200 {array} Book
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /api/v1/book [get]
func GetAllBooks(ctx *fiber.Ctx) {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	if err := ctx.JSON(books); err != nil {
		_ = ctx.Status(500).Error()
	}
}


// GetBook godoc
// @Summary Get Book by id
// @Description Get Book by id
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} Book
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /api/v1/book/{id} [get]
func GetBook(ctx *fiber.Ctx) {
	id := ctx.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book,id)

	if book.ID == 0 { // book not found
		_ = ctx.Status(404).Error()
		return
	}

	if err := ctx.JSON(book); err != nil {
		_ = ctx.Status(500).Error()
		return
	}
}

// AddBook godoc
// @Summary Add new Book
// @Description Add new Book
// @Accept json
// @Produce json
// @Param book body Book true "Add book"
// @Success 200 {object} Book
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /api/v1/book [post]
func AddBook(ctx *fiber.Ctx) {
	db := database.DBConn
	book := new(Book)
	if err := ctx.BodyParser(book); err != nil {
		_ = ctx.Status(500).Error()
		return
	}

	db.Create(&book)
	if err := ctx.JSON(book); err != nil {
		_ = ctx.Status(500).Error()
		return
	}
}


// DeleteBook godoc
// @Summary Delete Book by id
// @Description Delete Book by id
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} Book
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /api/v1/book/{id} [delete]
func DeleteBook(ctx *fiber.Ctx) {
	id := ctx.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book,id)

	if book.Title == "" {
		_ = ctx.Status(404).Error()
		return
	}
	db.Delete(&book)
	if err := ctx.JSON(book); err != nil {
		_ = ctx.Status(500).Error()
		return
		return
	}
}


// UpdateBook godoc
// @Summary update existing Book
// @Description update existing Book
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param book body Book true "Add book"
// @Success 200 {object} Book
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /api/v1/book/{id} [put]
func UpdateBook(ctx *fiber.Ctx){
	id := ctx.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book,id)

	if book.Title == "" {
		_ = ctx.Status(404).Error()
		return
	}
	newbook := new(Book)
	if err := ctx.BodyParser(newbook); err != nil {
		_ = ctx.Status(500).Error()
		return
	}

	book.Rating = newbook.Rating
	book.Title = newbook.Title
	book.Author = newbook.Author

	db.Save(&book)
	if err := ctx.JSON(book); err != nil {
		_ = ctx.Status(500).Error()
	}
}
