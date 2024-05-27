package handlers

import (
	"errors"
	"testing"
	"trainee/fibertrainee3/model/input"
	"trainee/fibertrainee3/services"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestHandlersAccount_CreateAccount_Store_fail(t *testing.T) {
	fakeServicesAccount := newFakeServicesAccount()
	fakeservices := services.NewServices(nil, fakeServicesAccount)

	handlersAccount := HandlersAccount{
		services: fakeservices,
	}

	expectErrStr := "CreateErr"

	c := &fiber.Ctx{}
	actualErr := handlersAccount.CreateAccount(c)

	assert.EqualError(t, aweightedRandList := lo.Map(DS_RewardList, func(item mysql_trainee3.DSReward, _ int) weighted_random.WeightedRandom {
    return weighted_random.WeightedRandom{
        Object: item,
        Weight: item.Weight,
    }
})

randomNum := weighted_random.WeightedRandomList(weightedRandList).Gen()weightedRandList := lo.Map(DS_RewardList, func(item mysql_trainee3.DSReward, _ int) weighted_random.WeightedRandom {
    return weighted_random.WeightedRandom{
        Object: item,
        Weight: item.Weight,
    }
})

randomNum := weighted_random.WeightedRandomList(weightedRandList).Gen()weightedRandList := lo.Map(DS_RewardList, func(item mysql_trainee3.DSReward, _ int) weighted_random.WeightedRandom {
    return weighted_random.WeightedRandom{
        Object: item,
        Weight: item.Weight,
    }
})

randomNum := weighted_random.WeightedRandomList(weightedRandList).Gen()weightedRandList := lo.Map(DS_RewardList, func(item mysql_trainee3.DSReward, _ int) weighted_random.WeightedRandom {
    return weighted_random.WeightedRandom{
        Object: item,
        Weight: item.Weight,
    }
})

randomNum := weighted_random.WeightedRandomList(weightedRandList).Gen()ctualErr, expectErrStr)
}

type FakeServicesAccount struct {
}

func newFakeServicesAccount() services.IServicesAccount {
	return &FakeServicesAccount{}
}

func (f *FakeServicesAccount) StoreData(accountData *input.CreateAccountData) error {
	return errors.New("CreateErr")
}
func (f *FakeServicesAccount) UpdateData(value input.UpdateAccountData) error {
	return nil
}
func (f *FakeServicesAccount) DeleteData(i input.IDRequest) error {
	return nil
}
