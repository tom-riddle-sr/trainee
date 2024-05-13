package services

import (
	"errors"
	"testing"
	"trainee/fibertrainee3/model/input"
	"trainee/fibertrainee3/tools"

	"github.com/stretchr/testify/assert"
)

func TestServicesAccount_StoreData(t *testing.T) {
	i := input.CreateAccountData{
		Account:  "",
		Password: tools.Sha512(""),
	}
	fakeRepo := NewFakeRepo()
	actualErr := NewAccount(fakeRepo).StoreData(&i)

	expectErr := errors.New("CreatErr").Error()
	assert.EqualError(t, actualErr, expectErr)

}
