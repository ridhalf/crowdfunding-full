package main

import (
	"crowdfunding/app"
	"crowdfunding/auth"
	"crowdfunding/controller"
	"crowdfunding/helper"
	"crowdfunding/middleware"
	"crowdfunding/repository"
	"crowdfunding/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	gin.SetMode(gin.DebugMode)
	app.Env()
	db := app.NewDB()

	//repositories
	userRepository := repository.NewUserRepository(db)
	campaignRepository := repository.NewCampaignRepository(db)
	transactionRepository := repository.NewTransactionRepository(db)
	//services
	userService := service.NewUserServiceImpl(userRepository)
	campaignService := service.NewCampaignService(campaignRepository)
	paymentService := service.NewPaymentService(transactionRepository, campaignRepository)
	transactionService := service.NewTransactionService(transactionRepository, campaignRepository, paymentService)
	//middleware
	authJwt := auth.NewJwtService()
	authMiddleware := middleware.AuthMiddleware(authJwt, userService)

	//controllers
	userController := controller.NewUserController(userService, authJwt)
	campaignController := controller.NewCampaignController(campaignService)
	transactionController := controller.NewTransactionController(transactionService, paymentService)
	router := gin.Default()
	//blocked by cors policy
	router.Use(cors.Default())
	//blocked by cors policy
	router.Static("/images", "./images")

	api := router.Group("/api/v1")
	api.POST("/users", userController.Register)
	api.POST("/users/login", userController.Login)
	api.POST("/users/email_checker", userController.IsEmailAvailable)
	api.POST("/users/avatar", authMiddleware, userController.UploadAvatar)
	api.GET("/users/fetch", authMiddleware, userController.FetchUser)

	api.GET("/campaigns", campaignController.FindAll)
	api.GET("/campaigns/:id", campaignController.FindByID)
	api.POST("/campaigns", authMiddleware, campaignController.Create)
	api.PUT("/campaigns/:id", authMiddleware, campaignController.Update)
	api.POST("/campaigns/image", authMiddleware, campaignController.UploadImage)

	api.GET("/campaigns/:id/transactions", authMiddleware, transactionController.FindByCampaignID)
	api.GET("/transactions", authMiddleware, transactionController.FindByUserID)
	api.POST("/transactions", authMiddleware, transactionController.Create)
	api.POST("/transactions/notification", transactionController.GetNotification)

	err := router.Run(os.Getenv("DOMAIN"))
	helper.PanicIfError(err)
}
