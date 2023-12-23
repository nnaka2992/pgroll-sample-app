package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func createUser(c echo.Context) error {
	user := generateUser()
	if _, err := pool.Exec(context.Background(), "INSERT INTO users (first_name, last_name, date_of_birth) VALUES ($1, $2, $3)", &user.FirstName, &user.LastName, &user.DateOfBirth); err != nil {
		return fmt.Errorf("Could not insert user: ", err.Error())
	}
	return c.String(http.StatusOK, "createUser")
}

func getUserByID(c echo.Context) error {
	id := c.Param("id")

	user := User{}
	row := pool.QueryRow(context.Background(), "SELECT * FROM users WHERE id = $1", id)
	if err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.DateOfBirth); err != nil {
		return fmt.Errorf("Could not retrieve columns: ", err.Error())
	}
	b, err := json.Marshal(user); if err != nil {
		return fmt.Errorf("Could not marshal user: ", err.Error())
	}

	return c.String(http.StatusOK, string(b) + "\n")
}
