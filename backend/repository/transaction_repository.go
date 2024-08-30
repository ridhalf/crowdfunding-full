package repository

import "crowdfunding/model/domain"

type TransactionRepository interface {
	FindByCampaignID(campaignID int) ([]domain.Transaction, error)
	FindByUserID(userID int) ([]domain.Transaction, error)
	Create(transaction domain.Transaction) (domain.Transaction, error)
	Update(transaction domain.Transaction) (domain.Transaction, error)
	FindByID(ID int) (domain.Transaction, error)
}
