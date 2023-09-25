package PGRepository

import (
	"database/sql"

	"github.com/cocoasterr/net_http/models"
)

type UserRepository struct {
	Repository
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		Repository{
			Db:    db,
			Model: &models.Users{},
		},
	}
}
