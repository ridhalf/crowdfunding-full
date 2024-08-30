package web

import (
	"crowdfunding/model/domain"
	"time"
)

type TransactionResponseCampaign struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Amount    int    `json:"amount"`
	CreatedAt string `json:"created_at"`
}

func ToTransactionResponseCampaign(transaction domain.Transaction) TransactionResponseCampaign {
	return TransactionResponseCampaign{
		ID:        transaction.ID,
		Name:      transaction.User.Name,
		Amount:    transaction.Amount,
		CreatedAt: transaction.CreatedAt.String(),
	}
}
func ToTransactionResponseCampaigns(transactions []domain.Transaction) []TransactionResponseCampaign {
	if len(transactions) == 0 {
		return []TransactionResponseCampaign{}
	}
	var transactionResponses []TransactionResponseCampaign
	for _, transaction := range transactions {
		transactionResponses = append(transactionResponses, ToTransactionResponseCampaign(transaction))
	}
	return transactionResponses
}

type TransactionResponseUser struct {
	ID        int                                   `json:"id"`
	Amount    int                                   `json:"amount"`
	Status    string                                `json:"status"`
	CreatedAt time.Time                             `json:"created_at"`
	Campaign  TransactionResponseUserDetailCampaign `json:"campaign"`
}
type TransactionResponseUserDetailCampaign struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func ToTransactionResponseUser(transaction domain.Transaction) TransactionResponseUser {
	var url string
	if len(transaction.Campaign.CampaignImages) > 0 {
		url = transaction.Campaign.CampaignImages[0].FileName
	}

	Campaign := TransactionResponseUserDetailCampaign{
		Name:     transaction.Campaign.Name,
		ImageURL: url,
	}
	return TransactionResponseUser{
		ID:        transaction.ID,
		Amount:    transaction.Amount,
		Status:    transaction.Status,
		CreatedAt: transaction.CreatedAt,
		Campaign:  Campaign,
	}
}
func ToTransactionResponseUsers(transactions []domain.Transaction) []TransactionResponseUser {
	if len(transactions) == 0 {
		return []TransactionResponseUser{}
	}
	var transactionResponses []TransactionResponseUser
	for _, transaction := range transactions {
		transactionResponses = append(transactionResponses, ToTransactionResponseUser(transaction))
	}
	return transactionResponses
}

type TransactionResponseCreate struct {
	ID         int    `json:"id"`
	CampaignID int    `json:"campaign_id"`
	UserID     int    `json:"user_id"`
	Amount     int    `json:"amount"`
	Status     string `json:"status"`
	Code       string `json:"code"`
	PaymentURL string `json:"payment_url"`
}

func ToTransactionResponseCreate(transaction domain.Transaction) TransactionResponseCreate {
	return TransactionResponseCreate{
		ID:         transaction.ID,
		CampaignID: transaction.CampaignID,
		UserID:     transaction.UserID,
		Amount:     transaction.Amount,
		Status:     transaction.Status,
		Code:       transaction.Code,
		PaymentURL: transaction.PaymentURL}
}
