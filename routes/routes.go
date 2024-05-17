package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pongsapak-suwa/REST-API-CRUD-Go/handler"
	"github.com/pongsapak-suwa/REST-API-CRUD-Go/utils"
)

func Routes(app *fiber.App) {
	auth := app.Group("/api")
    auth.Post("/register", handler.Register)
    auth.Post("/login", handler.Login)

    secured := auth.Group("/d")
    secured.Use(utils.AuthMiddleware())

    secured.Get("/books", handler.GetBooks)
    secured.Get("/books/:id", handler.GetBook)
    secured.Post("/books", handler.AddBook)
    secured.Put("/books/:id", handler.UpdateBook)
    secured.Delete("/books/:id", handler.DeleteBook)
}
