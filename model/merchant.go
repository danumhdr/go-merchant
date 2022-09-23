package model

type Merchant struct {
	Id              int    `json:"id"`
	MerchantId      int    `json:"merchant_id"`
	MerchantName    string `json:"merchant_name"`
	Bill            int    `json:"bill"`
	DateTransaction string `json:"date_transaction"`
}
