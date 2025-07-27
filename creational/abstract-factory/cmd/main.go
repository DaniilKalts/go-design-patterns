package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"

	"github.com/DaniilKalts/go-design-patterns/abstract-factory/internal/adapters/database"
	"github.com/DaniilKalts/go-design-patterns/abstract-factory/internal/domain"
	"github.com/DaniilKalts/go-design-patterns/abstract-factory/internal/repository"
)

func main() {
	db, err := database.InitMongoDB(
		"mongodb://mongo:27017",
		"abstract-factory",
	)
	if err != nil {
		log.Fatal(err)
	}

	repos := repository.NewRepoFactory(db)

	_, err = repos.UserRepo().CreateUser(
		context.Background(),
		&domain.User{
			ID:        uuid.New(),
			FirstName: "Daniil",
			LastName:  "Kalts",
			Email:     "gopher@gmail.com",
			Age:       19,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	user, err := repos.UserRepo().GetUserByEmail(
		context.Background(), "gopher@gmail.com",
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user)
}
