package model

type Outlet struct {
	Id              int    `json:"id"`
	MerchantId      int    `json:"merchant_id"`
	OutletId        int    `json:"outlet_id"`
	MerchantName    string `json:"merchant_name"`
	OutletName      string `json:"outlet_name"`
	Bill            int    `json:"bill"`
	DateTransaction string `json:"date_transaction"`
}
