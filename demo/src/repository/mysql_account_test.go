package repository_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
	accountRepo "repository"
	"testing"
)

func TestGetByName(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id", "name", "email"}).AddRow(1, "Stub Name 1", "Stun Email 2")
	sql := `SELECT id, name, email FROM account WHERE name=?`

	mock.ExpectQuery(sql).WillReturnRows(rows)
	ar := accountRepo.NewMySQLAccountRepository(db)
	account, err := ar.GetByName(context.TODO(), "Stub Name 1")
	assert.NoError(t, err)
	assert.NotNil(t, account)
}
