package service

import (
	"crowdfunding/model/domain"
	"crowdfunding/model/web"
)

type PaymentService interface {
	GetPaymentUrl(transaction domain.Payment, user domain.User) (string, error)
	ProcessPayment(request web.TransactionRequestNotification) error
}
