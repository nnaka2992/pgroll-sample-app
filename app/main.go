package main

import (
	"os"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	var err error
	pool, err = NewPoolFromEnv()
	if err != nil {
		log.Println("Unable to connect to database: ", err)
		os.Exit(1)
	}
	defer pool.Close()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users/:id", getUserByID)
	e.POST("/users/new", createUser)

	e.Logger.Fatal(e.Start(":8080"))
}
