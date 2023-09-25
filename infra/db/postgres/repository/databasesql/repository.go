package PGRepository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/cocoasterr/net_http/models"
)

type RepoInterface interface {
	Create(ctx context.Context, model interface{}, tbName string) error
	Index(ctx context.Context, page, limit int, tbName string) ([]interface{}, interface{}, error)
	FindBy(ctx context.Context, tbName, key string, value interface{}) (map[string]interface{}, error)
	Delete(ctx context.Context, tbname, key string, value interface{}) error
}
type Repository struct {
	Db    *sql.DB
	Model models.BaseModel
}

func (r *Repository) Create(ctx context.Context, model map[string]interface{}) error {
	tbName := r.Model.TbName()
	var key, value []string
	var values []interface{}
	var i int
	for k, v := range model {
		key = append(key, k)
		value = append(value, fmt.Sprintf("$%d", i+1))
		values = append(values, v)
		i++
	}
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tbName, strings.Join(key, ", "), strings.Join(value, ", "))

	trx, err := r.Db.Begin()
	if err != nil {
		return err
	}
	_, err = trx.ExecContext(ctx, query, values...)
	if err != nil {
		trx.Rollback()
		return err
	}
	err = trx.Commit()
	if err != nil {
		trx.Rollback()
		return err
	}
	return nil
}

func (r *Repository) Index(ctx context.Context, page, limit int) ([]map[string]interface{}, interface{}, error) {
	tbName := r.Model.TbName()
	offset := (page - 1) * limit
	query := fmt.Sprintf("SELECT * FROM %s LIMIT %d OFFSET %d", tbName, limit, offset)
	rows, _ := r.Db.QueryContext(ctx, query)
	if rows == nil {
		return nil, nil, nil
	}
	data, _ := getResIndex(rows)

	var total int
	query_total := fmt.Sprintf("SELECT COUNT(id) FROM %s", tbName)
	if err := r.Db.QueryRowContext(ctx, query_total).Scan(&total); err != nil {
		return nil, nil, err
	}

	return data, total, nil
}
func (r *Repository) FindBy(ctx context.Context, key string, value interface{}) ([]map[string]interface{}, error) {
	tbName := r.Model.TbName()

	typeValue := reflect.ValueOf(value)
	if typeValue.Kind() == reflect.String {
		value = fmt.Sprintf("'%s'", value)
	}
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s = %v", tbName, key, value)
	rows, _ := r.Db.QueryContext(ctx, query)
	if rows == nil {
		return nil, nil
	}
	data, _ := getResIndex(rows)

	return data, nil
}

func (r *Repository) FindCustomQuery(ctx context.Context, condition string) ([]map[string]interface{}, error) {
	tbName := r.Model.TbName()
	query := fmt.Sprintf("SELECT * FROM %s %s", tbName, condition)
	rows, _ := r.Db.QueryContext(ctx, query)
	if rows == nil {
		return nil, nil
	}
	data, _ := getResIndex(rows)

	return data, nil
}

func (r *Repository) Update(ctx context.Context, id string, model map[string]interface{}) error {
	tbName := r.Model.TbName()

	existingData, err := r.FindBy(ctx, "id", id)
	if err != nil {
		return err
	}
	if len(existingData) == 0 {
		return errors.New("data not found")
	}

	var updates []string
	var values []interface{}
	var i int
	for k, v := range model {
		updates = append(updates, fmt.Sprintf("%s=$%d", k, i+1))
		values = append(values, v)
		i++
	}
	values = append(values, id)
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", tbName, strings.Join(updates, ", "), i+1)

	trx, err := r.Db.Begin()
	if err != nil {
		return err
	}
	_, err = trx.ExecContext(ctx, query, values...)
	if err != nil {
		trx.Rollback()
		return err
	}
	err = trx.Commit()
	if err != nil {
		trx.Rollback()
		return err
	}
	return nil
}

func (r *Repository) Delete(ctx context.Context, id string) error {
	tbName := r.Model.TbName()

	existingData, err := r.FindBy(ctx, "id", id)
	if err != nil {
		return err
	}
	if len(existingData) == 0 {
		return errors.New("data not found")
	}
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", tbName)
	trx, err := r.Db.Begin()
	if err != nil {
		return err
	}
	_, err = trx.ExecContext(ctx, query, id)
	if err != nil {
		trx.Rollback()
		return err
	}
	err = trx.Commit()
	if err != nil {
		trx.Rollback()
		return err
	}
	return nil
}

func getResIndex(rows *sql.Rows) ([]map[string]interface{}, error) {
	lisColumn, _ := rows.Columns()
	var res []map[string]interface{}
	for rows.Next() {
		dest := make([]interface{}, len(lisColumn))
		for i := range lisColumn {
			dest[i] = new(interface{})
		}
		err := rows.Scan(dest...)
		if err != nil {
			return nil, err
		}
		itermap := make(map[string]interface{})
		for i, colName := range lisColumn {
			itermap[colName] = *dest[i].(*interface{})
		}
		res = append(res, itermap)
	}
	return res, nil
}
