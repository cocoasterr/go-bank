package services

import (
	PGRepository "github.com/cocoasterr/net_http/infra/db/postgres/repository/databasesql"
	"github.com/cocoasterr/net_http/models"
)

type BankService struct {
	BaseService
}

func NewAuthService(repo PGRepository.Repository) *BankService {
	return &BankService{
		BaseService: BaseService{
			Repo:  repo,
			Model: &models.Users{},
		},
	}
}

func (s *BankService) Transaction(userData map[string]interface{}, nominal int) map[string]interface{} {
	balance := int(userData["balance"].(int64)) + nominal
	if balance < 0 {
		return nil
	}
	userData["balance"] = balance + nominal

	return userData
}
