package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pongsapak-suwa/REST-API-CRUD-Go/database"
	"github.com/pongsapak-suwa/REST-API-CRUD-Go/models"
)

func GetBooks(c *fiber.Ctx) error {
    var books []models.Book
    database.DB.Find(&books)
    return c.JSON(books)
}

func GetBook(c *fiber.Ctx) error {
    id := c.Params("id")
    var book models.Book
    if err := database.DB.First(&book, id).Error; err != nil {
        return c.Status(404).JSON(fiber.Map{"error": "Record not found"})
    }
    return c.JSON(book)
}

func AddBook(c *fiber.Ctx) error {
    book := new(models.Book)

    if err := c.BodyParser(book); err != nil {
        return err
    }

    database.DB.Create(&book)
    return c.JSON(book)
}

func UpdateBook(c *fiber.Ctx) error {
    id := c.Params("id")
    var book models.Book
    if err := database.DB.First(&book, id).Error; err != nil {
        return c.Status(404).JSON(fiber.Map{"error": "Record not found"})
    }

    updatedBook := new(models.Book)
    if err := c.BodyParser(updatedBook); err != nil {
        return err
    }

    database.DB.Model(&book).Updates(updatedBook)
    return c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
    id := c.Params("id")
    var book models.Book
    if err := database.DB.First(&book, id).Error; err != nil {
        return c.Status(404).JSON(fiber.Map{"error": "Record not found"})
    }

    database.DB.Delete(&book)
    return c.SendStatus(fiber.StatusNoContent)
}
