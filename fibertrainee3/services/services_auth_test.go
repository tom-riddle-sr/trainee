package services

import (
	"database/sql"
	"errors"
	"testing"
	"trainee/fibertrainee3/model/input"
	"trainee/fibertrainee3/repository"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func Test_auth_Login_find_account_sql_no_row(t *testing.T) {
	i := input.CreateAccountData{
		Account:  "test",
		Password: "test",
	}

	expectErr := "查無此帳"

	fakeRepo := NewFakeRepo()

	serviceAuth := ServicesAuth{
		repo: fakeRepo,
	}

	_, actualErr := serviceAuth.Login(i)

	assert.EqualError(t, actualErr, expectErr)
}

type FakeRepo struct{}

func NewFakeRepo() repository.IRepo {
	return &FakeRepo{}
}

func (f *FakeRepo) Create(db *gorm.DB, value interface{}) error {
	return errors.New("CreateErr")
}

func (f *FakeRepo) Query(db *gorm.DB, query string, model interface{}, values ...interface{}) error {
	return sql.ErrNoRows
}

func (f *FakeRepo) Update(db *gorm.DB, condition string, value interface{}, model interface{}) error {
	return nil
}

func (f *FakeRepo) Delete(db *gorm.DB, condition string, value interface{}, model interface{}) error {
	return nil
}
