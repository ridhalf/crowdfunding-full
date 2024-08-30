package service

import (
	"crowdfunding/helper"
	"crowdfunding/model/domain"
	"crowdfunding/model/web"
	"crowdfunding/repository"
	"github.com/veritrans/go-midtrans"
	"os"
	"strconv"
	"strings"
)

type PaymentServiceImpl struct {
	transactionRepository repository.TransactionRepository
	campaignRepository    repository.CampaignRepository
}

func NewPaymentService(transactionRepository repository.TransactionRepository, campaignRepository repository.CampaignRepository) PaymentService {
	return &PaymentServiceImpl{
		transactionRepository: transactionRepository,
		campaignRepository:    campaignRepository,
	}
}

func (service PaymentServiceImpl) GetPaymentUrl(payment domain.Payment, user domain.User) (string, error) {
	midclient := midtrans.NewClient()
	midclient.ServerKey = os.Getenv("SERVER_KEY")
	midclient.ClientKey = os.Getenv("CLIENT_KEY")
	midclient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}
	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.Name,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  helper.ORDER_FORMAT + strconv.Itoa(payment.ID),
			GrossAmt: int64(payment.Amount),
		},
	}
	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}
	return snapTokenResp.RedirectURL, nil
}

func (service PaymentServiceImpl) ProcessPayment(request web.TransactionRequestNotification) error {
	transaction_id, _ := strconv.Atoi(strings.TrimPrefix(request.OrderID, "ORDER-"))
	transaction, err := service.transactionRepository.FindByID(transaction_id)
	if err != nil {
		return err
	}
	if request.PaymentType == helper.CREDIT_CARD && request.TransactionStatus == helper.CAPTURE && request.FraudStatus == helper.ACCEPT {
		transaction.Status = helper.PAID
	} else if request.TransactionStatus == helper.SETTLEMENT {
		transaction.Status = helper.PAID
	} else if request.TransactionStatus == helper.DENY || request.TransactionStatus == helper.EXPIRE || request.TransactionStatus == helper.CANCEL {
		transaction.Status = helper.CANCELLED
	}
	update, err := service.transactionRepository.Update(transaction)
	if err != nil {
		return err
	}
	campaign, err := service.campaignRepository.FindByID(update.CampaignID)
	if err != nil {
		return err
	}
	if update.Status == helper.PAID {
		campaign.BackerCount += 1
		campaign.CurrentAmount += update.Amount
		_, err = service.campaignRepository.Update(campaign)
		if err != nil {
			return err
		}
	}
	return nil
}
