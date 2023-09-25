package PGGormRepository

import "gorm.io/gorm"

type PersonService struct {
	PGGormRepository
}

func NewPersonService(db *gorm.DB) *PersonService {
	return &PersonService{
		PGGormRepository: PGGormRepository{
			DB: db,
		},
	}
}
