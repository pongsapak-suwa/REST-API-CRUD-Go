package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/pongsapak-suwa/REST-API-CRUD-Go/database"
	"github.com/pongsapak-suwa/REST-API-CRUD-Go/models"
	"github.com/pongsapak-suwa/REST-API-CRUD-Go/utils"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
    user := new(models.User)

    if err := c.BodyParser(user); err != nil {
        return err
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    user.Password = string(hashedPassword)
    database.DB.Create(&user)
    return c.SendStatus(fiber.StatusCreated)
}

func Login(c *fiber.Ctx) error {
    var inputUser models.User
    if err := c.BodyParser(&inputUser); err != nil {
        return err
    }

    var user models.User
    if err := database.DB.Where("username = ?", inputUser.Username).First(&user).Error; err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid username or password"})
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(inputUser.Password)); err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid username or password"})
    }

    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &utils.Claims{
        Username: user.Username,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(utils.JwtKey)
    if err != nil {
        return err
    }

    return c.JSON(fiber.Map{"token": tokenString})
}
