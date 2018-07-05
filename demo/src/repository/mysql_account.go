package repository

import (
	"context"
	"database/sql"
	"model"
)

type AccountRepository interface {
	GetByName(ctx context.Context, name string) (*model.Account, error)
}

type mysqlAccountRepository struct {
	Conn *sql.DB
}

func NewMySQLAccountRepository(Conn *sql.DB) AccountRepository {
	return &mysqlAccountRepository{Conn}
}

func (m *mysqlAccountRepository) GetByName(ctx context.Context, name string) (*model.Account, error) {
	sql := `SELECT id, name, email FROM account WHERE name=?`
	list, err := m.fetch(ctx, sql, name)
	if err != nil {
		return nil, err
	}

	a := &model.Account{}
	if len(list) > 0 {
		a = list[0]
	} else {
		return nil, model.NOT_FOUND_ERROR
	}

	return a, nil
}

func (m *mysqlAccountRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*model.Account, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make([]*model.Account, 0)
	for rows.Next() {
		t := new(model.Account)
		err = rows.Scan(
			&t.Id,
			&t.Name,
			&t.Email,
		)

		if err != nil {
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}
