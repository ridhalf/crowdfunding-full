package controller

import (
	"crowdfunding/helper"
	"crowdfunding/model/domain"
	"crowdfunding/model/web"
	"crowdfunding/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TransactionControllerImpl struct {
	transactionService service.TransactionService
	paymentService     service.PaymentService
}

func NewTransactionController(transactionService service.TransactionService, paymentService service.PaymentService) TransactionController {
	return &TransactionControllerImpl{
		transactionService: transactionService,
		paymentService:     paymentService,
	}
}

func (controller TransactionControllerImpl) FindByCampaignID(ctx *gin.Context) {
	var request web.TransactionRequestByCampaignID
	err := ctx.ShouldBindUri(&request)
	if err != nil {
		controller.failedTransactions(ctx, false)
		return
	}

	user := ctx.MustGet("user").(domain.User)
	request.User = user

	transactions, err := controller.transactionService.FindByCampaignID(request)
	if err != nil {
		controller.failedTransactions(ctx, true)
		return
	}
	response := web.ToTransactionResponseCampaigns(transactions)
	result := helper.Ok("list all transactions", response)
	ctx.JSON(http.StatusOK, result)
}
func (controller TransactionControllerImpl) FindByUserID(ctx *gin.Context) {
	user := ctx.MustGet("user").(domain.User)
	transactions, err := controller.transactionService.FindByUserID(user.ID)
	if err != nil {
		controller.failedTransactions(ctx, false)
		return
	}
	response := web.ToTransactionResponseUsers(transactions)
	result := helper.Ok("list all transactions", response)
	ctx.JSON(http.StatusOK, result)
}

func (controller TransactionControllerImpl) Create(ctx *gin.Context) {
	var transaction web.TransactionRequestCreate
	err := ctx.ShouldBindJSON(&transaction)
	if err != nil {
		result := helper.UnprocessableEntity("failed to create transaction", err)
		ctx.JSON(http.StatusUnprocessableEntity, result)
		return
	}
	user := ctx.MustGet("user").(domain.User)
	transaction.User = user
	create, err := controller.transactionService.Create(transaction)
	if err != nil {
		result := helper.UnprocessableEntity("failed to create campaign", err)
		ctx.JSON(http.StatusUnprocessableEntity, result)
	}
	response := web.ToTransactionResponseCreate(create)
	result := helper.Ok("save campaign", response)
	ctx.JSON(http.StatusOK, result)
}

func (controller TransactionControllerImpl) GetNotification(ctx *gin.Context) {
	var request web.TransactionRequestNotification
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		result := helper.BadRequest("failed to process notification", err)
		ctx.JSON(http.StatusBadRequest, result)
		return
	}
	err = controller.paymentService.ProcessPayment(request)
	if err != nil {
		result := helper.BadRequest("failed to process notification", err)
		ctx.JSON(http.StatusBadRequest, result)
		return
	}
	result := helper.Ok("successfully notification", nil)
	ctx.JSON(http.StatusOK, result)
}

func (controller TransactionControllerImpl) failedTransactions(ctx *gin.Context, forbidden bool) {
	if forbidden {
		response := helper.Forbidden("user is not the owner of the campaign", nil)
		ctx.JSON(http.StatusForbidden, response)
	} else {
		response := helper.BadRequest("failed to get campaign transaction", nil)
		ctx.JSON(http.StatusBadRequest, response)
	}

}
