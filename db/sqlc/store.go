package db

import (
	"database/sql"
)

type SQLStore struct {
	*Queries
	db *sql.DB
}

// Store provides all functions to excute db queries and transaction
type Store interface {
	Querier
}

// NewStore creates a new store
func NewStore(db *sql.DB) *SQLStore {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}
