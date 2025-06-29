package constructor

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type DBClient interface {
	Query(query string, args ...any) (*sql.Rows, error)
	// … other methods …
}

type postgresClient struct {
	db *sql.DB
}

func NewPostgresClient(dsn string, maxOpen, maxIdle int) (DBClient, error) {
	if dsn == "" {
		return nil, fmt.Errorf("dsn cannot be empty")
	}
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(maxOpen)
	db.SetMaxIdleConns(maxIdle)
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &postgresClient{db: db}, nil
}

func (c *postgresClient) Query(q string, args ...any) (*sql.Rows, error) {
	return c.db.Query(q, args...)
}

func RunDBClientDemo() {
	client, err := NewPostgresClient(os.Getenv("DATABASE_URL"), 20, 5)
	if err != nil {
		fmt.Println("Error creating DB client:", err)
		return
	}
	defer func() {
		if db, ok := client.(*postgresClient); ok {
			db.db.Close()
		}
	}()
	fmt.Printf("DB client created successfully: %T\n", client)
}
