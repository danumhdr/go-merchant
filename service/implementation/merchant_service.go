package implementation

import (
	"go-merchant/model"
	"go-merchant/repository"
	"go-merchant/repository/implementation"
	"go-merchant/service"
	"log"
)

type MerchantServe struct {
	Repository repository.MerchantRepository
}

func NewMerchantServe(merchanRepository repository.MerchantRepository) service.MerchantService {
	return &MerchantServe{Repository: merchanRepository}
}

func (merchantServe *MerchantServe) GetReportMerchant(userId string, month int, limit int, page int) ([]model.Merchant, error) {
	var result []model.Merchant
	result, err := implementation.NewMerchantRepo().GetReportMerchant(userId, month, limit, page)
	if err != nil {
		log.Println(err.Error())
		return result, err
	}
	return result, nil
}

func (merchantServe *MerchantServe) GetReportOutlet(userId string, month int, limit int, page int) ([]model.Outlet, error) {
	var result []model.Outlet
	result, err := implementation.NewMerchantRepo().GetReportOutlet(userId, month, limit, page)
	if err != nil {
		log.Println(err.Error())
		return result, err
	}
	return result, nil
}
