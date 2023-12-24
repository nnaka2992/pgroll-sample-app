package main

import (
	"log"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/jackc/pgx/pgtype"
)

type User struct {
	Id          pgtype.Int8
	FirstName   pgtype.Varchar
	LastName    pgtype.Varchar
	DateOfBirth pgtype.Date
}

func generateUser() User {
	date, err := time.Parse("2006-01-02", faker.Date())
	if err != nil {
		log.Println("Unable to parse date: ", err)
	}
	user := User{
		FirstName:   pgtype.Varchar{String: faker.FirstName(), Status: pgtype.Present},
		LastName:    pgtype.Varchar{String: faker.LastName(), Status: pgtype.Present},
		DateOfBirth: pgtype.Date{Time: date, Status: pgtype.Present},
	}
	return user
}
