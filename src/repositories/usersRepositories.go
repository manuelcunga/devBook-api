package repositories

import (
	"book-api/src/entities"
	"database/sql"
)

// restes
type users struct {
	db *sql.DB
}

// teste
func NewUserRepository(db *sql.DB) *users {
	return &users{db}
}

func (u users) Create(user entities.User) (uint64, error) {
	return 0, nil
}
