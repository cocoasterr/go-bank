package PGRepository

import (
	"database/sql"

	"github.com/cocoasterr/net_http/models"
)

type MutationRepository struct {
	Repository
}

func NewMutationRepository(db *sql.DB) *MutationRepository {
	return &MutationRepository{
		Repository{
			Db:    db,
			Model: &models.Mutation{},
		},
	}
}
