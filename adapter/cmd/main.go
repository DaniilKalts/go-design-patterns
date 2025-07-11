package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/daniilkalts/go-design-patterns/adapter/internal/domain/person"
)

func main() {
	dsn := "host=postgres user=gopher password=gopher dbname=adapter port=5432 sslmode=disable TimeZone=UTC"

	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)
	if err != nil {
		panic("Hmm, i thought it must work...")
	}

	if err := db.AutoMigrate(&person.Person{}); err != nil {
		panic("Hmm, i thought it must work...")
	}

	// === Client ===
	repo := person.NewPersonRepository(db)

	me := &person.Person{
		FirstName: "Daniil",
		LastName:  "Kalts",
	}

	if err := repo.Create(me); err != nil {
		panic("Hmm, i thought it must work...")
	}

	persons, err := repo.GetAll()
	if err != nil {
		panic("Hmm, i thought it must work...")
	}

	fmt.Println(persons)
	// ==============
}
