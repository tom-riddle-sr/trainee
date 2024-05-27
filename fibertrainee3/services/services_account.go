package services

import (
	"trainee/fibertrainee3/database/mysql"
	"trainee/fibertrainee3/model/entity/mysql/fibertrainee"
	"trainee/fibertrainee3/model/input"
	"trainee/fibertrainee3/repository"
	"trainee/fibertrainee3/tools"
)

type IServicesAccount interface {
	StoreData(accountData *input.CreateAccountData) error
	UpdateData(value input.UpdateAccountData) error
	DeleteData(i input.IDRequest) error
}

type ServicesAccount struct {
	repo *repository.Repo
}

func NewAccount(repo *repository.Repo) IServicesAccount {
	return &ServicesAccount{
		repo: repo,
	}
}

func (account *ServicesAccount) StoreData(accountData *input.CreateAccountData) error {
	dbAccountData := &fibertrainee.AccountDatum{
		Account:  accountData.Account,
		Password: accountData.Password,
	}
	accountData.Password = tools.Sha512(accountData.Password)
	if err := account.repo.SqlRepo.Create(mysql.GetDB(), dbAccountData); err != nil {
		return err
	}
	return nil
}

func (account *ServicesAccount) UpdateData(value input.UpdateAccountData) error {
	accountdata := fibertrainee.AccountDatum{}
	if err := account.repo.SqlRepo.Query(mysql.GetDB(), "id = ? and account = ?", &accountdata, value.ID, value.Account); err != nil {
		return err
	}
	if err := account.repo.SqlRepo.Update(mysql.GetDB(), "id = ?", value.ID, &fibertrainee.AccountDatum{
		Account:  accountdata.Account,
		Password: tools.Sha512(accountdata.Password),
		ID:       accountdata.ID,
	}); err != nil {
		return err
	}
	return nil
}

func (account *ServicesAccount) DeleteData(i input.IDRequest) error {
	accountdata := fibertrainee.AccountDatum{}
	if err := account.repo.SqlRepo.Query(mysql.GetDB(), "id = ?", i.ID, &accountdata); err != nil {
		return err
	}

	return account.repo.SqlRepo.Delete(mysql.GetDB(), "id= ?", i.ID, &accountdata)
}
