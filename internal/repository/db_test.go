package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDBConnection(t *testing.T) {
	_, db, _ := NewDBConnection(&Config{
		Host:     "localhost",
		Port:     5432,
		User:     "root",
		Password: "secret",
		DBName:   "alerts",
	})
	assert.NotNil(t, db, "Database connection should not be nil")
}
