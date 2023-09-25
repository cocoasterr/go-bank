package services

import (
	"context"

	PGRepository "github.com/cocoasterr/net_http/infra/db/postgres/repository/databasesql"
	"github.com/cocoasterr/net_http/models"
)

type BaseServiceInterface interface {
	CreateService(ctx context.Context, payload map[string]interface{}) error
	IndexService(ctx context.Context, page, limit int) ([]map[string]interface{}, interface{}, error)
	FindByService(ctx context.Context, key string, value interface{}) (map[string]interface{}, error)
	UpdateService(ctx context.Context, payload map[string]interface{}, id string) error
	DeleteProduct(ctx context.Context, id string) error
}

type BaseService struct {
	Repo  PGRepository.Repository
	Model models.BaseModel
}

func NewService(productRepo PGRepository.Repository) *BaseService {
	return &BaseService{
		Repo: productRepo,
	}
}

func (s *BaseService) CreateService(ctx context.Context, payload map[string]interface{}) error {
	createPayload := s.Model.ModelCreate(payload)
	return s.Repo.Create(ctx, createPayload)
}

func (s *BaseService) IndexService(ctx context.Context, page, limit int) ([]map[string]interface{}, interface{}, error) {
	datares, total, err := s.Repo.Index(ctx, page, limit)

	if err != nil {
		return nil, nil, err
	}

	return datares, total, nil
}

func (s *BaseService) FindByService(ctx context.Context, key string, value interface{}) (map[string]interface{}, error) {
	datares, err := s.Repo.FindBy(ctx, key, value)
	if err != nil {
		return nil, err
	}
	if len(datares) != 1 {
		return nil, err
	}
	return datares[0], nil
}
func (s *BaseService) FindByCustomConditionService(ctx context.Context, condition string) ([]map[string]interface{}, error) {
	datares, err := s.Repo.FindCustomQuery(ctx, condition)
	if err != nil {
		return nil, err
	}

	return datares, nil
}

func (s *BaseService) UpdateService(ctx context.Context, payload map[string]interface{}, id string) error {
	updatePayload := s.Model.ModelUpdate(payload)
	return s.Repo.Update(ctx, id, updatePayload)
}

func (s *BaseService) DeleteService(ctx context.Context, id string) error {
	err := s.Repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
