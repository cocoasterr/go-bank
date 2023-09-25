package PGGormRepository

import (
	"fmt"
	"log"
	"reflect"

	"gorm.io/gorm"
)

type PGGormRepository struct {
	DB *gorm.DB
}

func NewPGGormRepository(db *gorm.DB) *PGGormRepository {
	return &PGGormRepository{
		DB: db,
	}
}

func (r *PGGormRepository) Create(model interface{}) error {
	trx := r.DB.Begin()
	defer trx.Commit()

	if err := trx.Create(model).Error; err != nil {
		log.Fatal("Rolling Back")
		trx.Rollback()
		return err
	}
	return nil
}

// offset := (page - 1) * limit
// res := r.DB.Find(model).Limit(limit).Offset(offset)
func (r *PGGormRepository) Index(tbName string, page, limit int) ([]map[interface{}]interface{}, error) {
	offset := (page - 1) * limit
	query := fmt.Sprintf("SELECT * FROM %s OFFSET %d LIMIT %d", tbName, offset, limit)
	var result []map[interface{}]interface{}
	if err := r.DB.Raw(query).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *PGGormRepository) FindBy(tbName, key string, value any) (map[interface{}]interface{}, error) {
	valueType := reflect.TypeOf(value)
	if valueType == reflect.TypeOf("") {
		value = fmt.Sprintf("'%s'", value)
	}
	query := fmt.Sprintf("SELECT * FROM %s where %s = %v", tbName, key, value)
	var result map[interface{}]interface{}

	res := r.DB.Raw(query).Scan(&result)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, res.Error
	}
	return result, nil
}

func (r *PGGormRepository) Update(model interface{}, id, tbName string) error {
	_, err := r.FindBy(tbName, "id", id)
	if err != nil {
		return err
	}
	if err = r.DB.Save(model).Error; err != nil {
		return err
	}
	return nil
}

func (r *PGGormRepository) Delete(model interface{}, id string) error {
	if err := r.DB.Delete(model).Where("id=?", id).Error; err != nil {
		return err
	}
	return nil
}
