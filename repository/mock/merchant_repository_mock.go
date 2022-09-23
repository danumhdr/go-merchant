package mock

import (
	"go-merchant/model"

	"github.com/stretchr/testify/mock"
)

type MerchantRepositoryMock struct {
	Mock mock.Mock
}

func (repository *MerchantRepositoryMock) GetReportMerchant(userId string, month int, limit int, page int) ([]model.Merchant, error) {
	var result []model.Merchant
	args := repository.Mock.Called(userId, month, limit, page)
	if args.Get(0) == nil {
		return result, args.Error(0)
	}
	result = args.Get(0).([]model.Merchant)
	return result, nil

}

func (repository *MerchantRepositoryMock) GetReportOutlet(userId string, month int, limit int, page int) ([]model.Outlet, error) {
	var result []model.Outlet
	args := repository.Mock.Called(userId, month, limit, page)
	if args.Get(0) == nil {
		return result, args.Error(0)
	}
	result = args.Get(0).([]model.Outlet)
	return result, nil
}
