package service

import "go-merchant/model"

type MerchantService interface {
	GetReportMerchant(userId string, month int, limit int, page int) ([]model.Merchant, error)
	GetReportOutlet(userId string, month int, limit int, page int) ([]model.Outlet, error)
}
