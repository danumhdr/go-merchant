package repository

import "go-merchant/model"

type MerchantRepository interface {
	GetReportMerchant(userId string, month int, limit int, page int) ([]model.Merchant, error)
	GetReportOutlet(userId string, month int, limit int, page int) ([]model.Outlet, error)
}
