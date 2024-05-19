package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/pongsapak-suwa/REST-API-CRUD-Go/database"
	"github.com/pongsapak-suwa/REST-API-CRUD-Go/models"
	"gorm.io/gorm"
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
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).SendString("Invalid book ID")
    }

    var book models.Book

    bookUpdate := new(models.Book)    
    if err := c.BodyParser(bookUpdate); err != nil {
        return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
    }

    if err := database.DB.First(&book, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return c.Status(fiber.StatusNotFound).SendString("Book not found")
        }
        return c.Status(fiber.StatusInternalServerError).SendString("Database error")
    }

    if bookUpdate.Title != "" {
        book.Title = bookUpdate.Title
    }
    if bookUpdate.Author != "" {
        book.Author = bookUpdate.Author
    }

    if err := database.DB.Save(&book).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).SendString("Failed to update book")
    }

    return c.Status(fiber.StatusOK).JSON(book)
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
