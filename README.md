# Go Fiber REST API with User Authentication

![image](https://github.com/pongsapak-suwa/REST-API-CRUD-Go/assets/94011063/330730f4-37f6-46df-8c6c-50232e20d054)

This is a simple REST API built with Go Fiber framework, Gorm ORM, and JWT for user authentication.

## Features

- User registration with password hashing
- User login with JWT authentication
- CRUD operations for managing books

## Prerequisites

Before running the application, make sure you have the following installed:

- Go (1.16 or higher)
- MySQL
- Go Fiber framework (`github.com/gofiber/fiber/v2`)
- Gorm ORM (`gorm.io/gorm`)
- JWT library (`github.com/golang-jwt/jwt`)

## Setup

1. Clone the repository:

    ```bash
    git clone https://github.com/pongsapak-suwa/REST-API-CRUD-Go.git
    cd REST-API-CRUD-Go
    ```

2. Install dependencies:

    ```bash
    go mod tidy
    ```

3. Setup MySQL database:

    - Create a new database.
    - Update the database connection details in `database/Connect()` function.

4. Run the application:

    ```bash
    go run .
    ```
