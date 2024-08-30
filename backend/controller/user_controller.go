package controller

import "github.com/gin-gonic/gin"

type UserController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	IsEmailAvailable(ctx *gin.Context)
	UploadAvatar(ctx *gin.Context)
	FetchUser(ctx *gin.Context)
}
