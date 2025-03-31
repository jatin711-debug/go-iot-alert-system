package repository

import (
	sqlc "alerts/db/sqlc" // SQLC-generated code
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// NewDBConnection establishes and returns a connection to the database
func NewDBConnection(dbConfig *Config) (*sqlc.Queries, *sql.DB, error) {
	// Build the connection string for PostgreSQL (example for local development)
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBName)

	// Open the database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open database: %v", err)
	}

	// Check if the database is reachable
	if err := db.Ping(); err != nil {
		return nil, nil, fmt.Errorf("failed to ping database: %v", err)
	}

	// Use SQLC to create a Queries object to run type-safe queries
	queries := sqlc.New(db)

	return queries, db, nil
}

// Config holds the database connection configuration
type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}
