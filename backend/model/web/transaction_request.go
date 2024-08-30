package web

import (
	"crowdfunding/model/domain"
)

type TransactionRequestByCampaignID struct {
	CampaignID int         `uri:"id" binding:"required"`
	User       domain.User `json:"user"`
}
type TransactionRequestCreate struct {
	Amount     int         `json:"amount" binding:"required"`
	CampaignID int         `json:"campaign_id" binding:"required"`
	User       domain.User `json:"user"`
}

type TransactionRequestNotification struct {
	TransactionStatus string `json:"transaction_status" binding:"required"`
	OrderID           string `json:"order_id" binding:"required"`
	PaymentType       string `json:"payment_type" binding:"required"`
	FraudStatus       string `json:"fraud_status" binding:"required"`
}
