package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pongsapak-suwa/REST-API-CRUD-Go/database"
	"github.com/pongsapak-suwa/REST-API-CRUD-Go/routes"
)


func main() {
    app := fiber.New()

    database.Connect()

    routes.Routes(app)

    err := app.Listen(":3000")
	if err != nil {
        panic(err)
    }
}
