package controller

import "github.com/gin-gonic/gin"

type TransactionController interface {
	FindByCampaignID(ctx *gin.Context)
	FindByUserID(ctx *gin.Context)
	Create(ctx *gin.Context)
	GetNotification(ctx *gin.Context)
}
