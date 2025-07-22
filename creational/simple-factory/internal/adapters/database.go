package adapters

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type DBExplorer interface {
	Ping() error
	Version() (string, error)
}

type explorer struct {
	db   *sql.DB
	kind string
}

func NewDBExplorer(dsn, kind string) (DBExplorer, error) {
	switch kind {
	case "postgres", "mysql":
		db, err := sql.Open(kind, dsn)
		if err != nil {
			return nil, err
		}
		return &explorer{db: db, kind: kind}, nil
	default:
		return nil, fmt.Errorf("unsupported database kind %s", kind)
	}
}

func (e *explorer) Ping() error {
	return e.db.Ping()
}

func (e *explorer) Version() (string, error) {
	var query string

	switch e.kind {
	case "postgres":
		query = `SELECT version();`
	case "mysql":
		query = `SELECT version();`
	default:
		return "", fmt.Errorf("unsupported database kind")
	}

	var version string
	if err := e.db.QueryRow(query).Scan(&version); err != nil {
		return "", err
	}

	return version, nil
}
