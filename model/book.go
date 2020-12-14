package model

import (
	"free-blog/database"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)


type Book struct {
	gorm.Model
	Title  string `json:"name"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) error {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	return c.JSON(books)
}
func GetBook(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var book Book
	db.Find(&book, id)
	return c.JSON(book)
}
func NewBook(c *fiber.Ctx) error {
	db := database.DBConn
	var book = new(Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(503).SendString("JSON Parse Error")
	}
	db.Create(&book)
	return c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {

	db := database.DBConn
	id := c.Params("id")
	var book Book
	db.Find(&book, id)
	if book.Title == "" {
		return c.Status(500).SendString("NO book found with ID ")
	}
	db.Delete(&book)
	return c.SendString("Book Successfully Deleted")

}
