package services

import (
	"context"
	"errors"
	"fmt"

	PGRepository "github.com/cocoasterr/net_http/infra/db/postgres/repository/databasesql"
	"github.com/cocoasterr/net_http/models"
)

type Mutation struct {
	BaseService
}

func NewMutationService(mutation PGRepository.Repository) *Mutation {
	return &Mutation{
		BaseService{
			Repo:  mutation,
			Model: &models.Mutation{},
		},
	}
}

func (s *Mutation) IndexMutation(ctx context.Context, account_number string) ([]map[string]interface{}, error) {
	conditionQuery := fmt.Sprintf("WHERE account_number = '%s' ORDER BY transaction_at DESC", account_number)
	data, err := s.Repo.FindCustomQuery(ctx, conditionQuery)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return data, nil
}
