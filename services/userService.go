package services

import "github.com/go-api/models"

func GetUsers() []models.User {
	users := []models.User{
		{
			Id:    "1",
			Name:  "new person",
			Email: "testemail",
		},
	}

	return users

}
