package main

import (
	"fmt"
	"time"

	"github.com/daniilkalts/go-design-patterns/simple-factory/internal/adapters"
)

func main() {
	postgresDSN := "postgres://gopher:gopher@postgres:5432/simple_factory?sslmode=disable"

	pgExplorer, err := adapters.NewDBExplorer(postgresDSN, "postgres")
	if err != nil {
		panic(err)
	}

	fmt.Println("Trying to connect to postgres...")
	if err := pgExplorer.Ping(); err != nil {
		panic(err)
	}

	pgVersion, err := pgExplorer.Version()
	if err != nil {
		fmt.Printf("Error getting postgres version: %s\n", err)
	}
	fmt.Printf("Postgres version: %s\n", pgVersion)

	time.Sleep(2 * time.Second)

	mysqlDSN := "gopher:gopher@tcp(mysql:3306)/simple_factory?parseTime=true&loc=Local"

	myExplorer, err := adapters.NewDBExplorer(mysqlDSN, "mysql")
	if err != nil {
		panic(err)
	}

	fmt.Println("Trying to connect to mysql...")
	if err := myExplorer.Ping(); err != nil {
		panic(err)
	}

	myVersion, err := myExplorer.Version()
	if err != nil {
		fmt.Printf("Error getting MySQL version: %s\n", err)
	}
	fmt.Printf("MySQL version: %s\n", myVersion)
}
