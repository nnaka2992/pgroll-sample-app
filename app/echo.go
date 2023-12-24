package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/jackc/pgx/pgtype"
)

func createUser(c echo.Context) error {
	user := generateUser()
	row := pool.QueryRow(context.Background(), "INSERT INTO users (first_name, last_name, date_of_birth, email) VALUES ($1, $2, $3, $4) RETURNING id;", &user.FirstName, &user.LastName, &user.DateOfBirth, &user.Email)
	if row == nil {
		return fmt.Errorf("Could not insert user")
	}
	var userId pgtype.Int8
	if err := row.Scan(&userId); err != nil {
		return fmt.Errorf("Could not retrieve id: ", err.Error())
	}
	return c.String(http.StatusOK, fmt.Sprintf("User with ID: %d is created\n", userId.Int))
}

func getUserByID(c echo.Context) error {
	id := c.Param("id")

	user := User{}
	row := pool.QueryRow(context.Background(), "SELECT id, first_name, last_name, date_of_birth, email FROM users WHERE id = $1", id)
	if err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.DateOfBirth, &user.Email); err != nil {
		return fmt.Errorf("Could not retrieve columns: ", err.Error())
	}
	b, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("Could not marshal user: ", err.Error())
	}

	return c.String(http.StatusOK, string(b)+"\n")
}
