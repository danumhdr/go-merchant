package implementation

import (
	"go-merchant/database"
	"go-merchant/model"
	"go-merchant/repository"
	"log"
)

type MerchantRepo struct {
}

func NewMerchantRepo() repository.MerchantRepository {
	return &MerchantRepo{}
}

func (merchantRepo *MerchantRepo) GetReportMerchant(userId string, month int, limit int, page int) ([]model.Merchant, error) {
	var result []model.Merchant
	offset := (page - 1) * limit
	record := database.DB.Raw("SELECT DISTINCT a.merchant_id,merchant_name,SUM(bill_total) as bill, "+
		"DATE(a.created_at) as date_transaction FROM transactions a "+
		"join merchants b on a.merchant_id = b.id "+
		"where b.user_id = ? and MONTH(a.created_at) = ? "+
		"group by a.merchant_id,DATE(a.created_at) "+
		"order by merchant_name,DATE(a.created_at) asc LIMIT ? OFFSET ?", userId, month, limit, offset).Scan(&result)
	if record.Error != nil {
		log.Println(record.Error)
		return result, record.Error
	}
	return result, nil
}

func (merchantRepo *MerchantRepo) GetReportOutlet(userId string, month int, limit int, page int) ([]model.Outlet, error) {
	var result []model.Outlet
	offset := (page - 1) * limit
	record := database.DB.Raw("SELECT DISTINCT a.merchant_id,merchant_name,a.outlet_id,c.outlet_name,SUM(bill_total) as bill, "+
		"DATE(a.created_at) as date_transaction FROM transactions a "+
		"join merchants b on a.merchant_id = b.id "+
		"join outlets c on a.outlet_id = c.id "+
		"where b.user_id = ? and MONTH(a.created_at) = ? "+
		"group by a.merchant_id,a.outlet_id,DATE(a.created_at) "+
		"order by c.outlet_name,DATE(a.created_at) asc LIMIT ? OFFSET ?", userId, month, limit, offset).Scan(&result)
	if record.Error != nil {
		log.Println(record.Error)
		return result, record.Error
	}
	return result, nil
}
