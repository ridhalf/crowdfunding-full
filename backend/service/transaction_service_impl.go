package service

import (
	"crowdfunding/helper"
	"crowdfunding/model/domain"
	"crowdfunding/model/web"
	"crowdfunding/repository"
	"errors"
	"strconv"
)

type TransactionServiceImpl struct {
	transactionRepository repository.TransactionRepository
	campaignRepository    repository.CampaignRepository
	paymentService        PaymentService
}

func NewTransactionService(transactionRepository repository.TransactionRepository, campaignRepository repository.CampaignRepository, paymentService PaymentService) TransactionService {
	return &TransactionServiceImpl{
		transactionRepository: transactionRepository,
		campaignRepository:    campaignRepository,
		paymentService:        paymentService,
	}
}

func (service TransactionServiceImpl) FindByCampaignID(request web.TransactionRequestByCampaignID) ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	campaign, err := service.campaignRepository.FindByID(request.CampaignID)
	if err != nil {
		return helper.ResultOrError(transactions, err)
	}

	if campaign.UserID != request.User.ID {
		return helper.ResultOrError(transactions, errors.New("user id is not match"))
	}

	transactions, err = service.transactionRepository.FindByCampaignID(request.CampaignID)
	if err != nil {
		return helper.ResultOrError(transactions, err)
	}
	return transactions, nil
}

func (service TransactionServiceImpl) FindByUserID(userID int) ([]domain.Transaction, error) {
	transactions, err := service.transactionRepository.FindByUserID(userID)
	if err != nil {
		return helper.ResultOrError(transactions, err)
	}
	return transactions, nil
}

func (service TransactionServiceImpl) Create(request web.TransactionRequestCreate) (domain.Transaction, error) {
	transaction := domain.Transaction{
		CampaignID: request.CampaignID,
		Amount:     request.Amount,
		UserID:     request.User.ID,
		Status:     helper.PENDING,
	}
	create, err := service.transactionRepository.Create(transaction)
	if err != nil {
		return helper.ResultOrError(create, err)
	}
	payment := domain.Payment{
		ID:     create.ID,
		Amount: create.Amount,
	}
	url, err := service.paymentService.GetPaymentUrl(payment, request.User)
	if err != nil {
		return create, err
	}
	create.PaymentURL = url
	create.Code = helper.ORDER_FORMAT + strconv.Itoa(create.ID)
	update, err := service.transactionRepository.Update(create)
	if err != nil {
		return helper.ResultOrError(transaction, err)
	}
	return update, nil
}
