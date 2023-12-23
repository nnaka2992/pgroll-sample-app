package main

import (
	"time"
	"log"

	"github.com/go-faker/faker/v4"
)

type User struct {
	Id int64
	FirstName string
	LastName string
	DateOfBirth time.Time
}

func generateUser() User {
	date, err := time.Parse("2006-01-02", faker.Date()); if err != nil {
		log.Println("Unable to parse date: ", err)
	}
	user := User{
		FirstName: faker.FirstName(),
		LastName: faker.LastName(),
		DateOfBirth: date,
	}
	return user
}
