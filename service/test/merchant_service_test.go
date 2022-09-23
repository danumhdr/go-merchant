package test

import (
	"go-merchant/model"
	"go-merchant/repository/mock"
	"go-merchant/service/implementation"
	"testing"

	"github.com/stretchr/testify/assert"
	mocks "github.com/stretchr/testify/mock"
)

var merchantRepository = &mock.MerchantRepositoryMock{Mock: mocks.Mock{}}
var merchantService = implementation.MerchantServe{Repository: merchantRepository}

/*
Id              int    `json:"id"`

	MerchantId      int    `json:"merchant_id"`
	MerchantName    string `json:"merchant_name"`
	Bill            int    `json:"bill"`
	DateTransaction string `json:"date_transaction"`
*/
func TestCategoryServiceSuccess(t *testing.T) {
	merchant := []model.Merchant{
		{
			Id:              1,
			MerchantId:      1,
			MerchantName:    "satu",
			Bill:            1000,
			DateTransaction: "2022-09-23",
		},
	}
	merchantRepository.Mock.On("GetReportMerchant", "1", 1, 1, 1).Return(merchant)

	assert.Equal(t, 1, merchant[0].Id)
}
