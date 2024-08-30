package controller

import (
	"crowdfunding/helper"
	"crowdfunding/model/domain"
	"crowdfunding/model/web"
	"crowdfunding/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CampaignControllerImpl struct {
	campaignService service.CampaignService
}

func NewCampaignController(campaignService service.CampaignService) CampaignController {
	return &CampaignControllerImpl{
		campaignService: campaignService,
	}
}

func (controller CampaignControllerImpl) FindAll(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Query("user_id"))
	campaigns, err := controller.campaignService.FindAll(userID)
	if err != nil {
		controller.failedGetCampaigns(ctx, false)
		return
	}
	result := helper.Ok("list all campaigns", web.ToCampaignsResponse(campaigns))
	ctx.JSON(http.StatusOK, result)
}

func (controller CampaignControllerImpl) FindByID(ctx *gin.Context) {
	var request web.CampaignRequestByID
	err := ctx.ShouldBindUri(&request)
	if err != nil {
		controller.failedGetCampaigns(ctx, false)
		return
	}
	campaign, err := controller.campaignService.FindByID(request)
	if err != nil {
		controller.failedGetCampaigns(ctx, false)
		return
	}
	response := web.ToCampaignDetailResponse(campaign)
	result := helper.Ok("campaign detail", response)
	ctx.JSON(http.StatusOK, result)
}

func (controller CampaignControllerImpl) Create(ctx *gin.Context) {
	requestCampaign := web.CampaignRequestCreate{}
	err := ctx.ShouldBindJSON(&requestCampaign)
	if err != nil {
		result := helper.UnprocessableEntity("failed to create campaign", err)
		ctx.JSON(http.StatusUnprocessableEntity, result)
		return
	}
	user := ctx.MustGet("user").(domain.User)
	requestCampaign.User = user

	create, err := controller.campaignService.Create(requestCampaign)
	if err != nil {
		result := helper.UnprocessableEntity("failed to create campaign", err)
		ctx.JSON(http.StatusUnprocessableEntity, result)
	}
	response := web.ToCampaignResponse(create)
	result := helper.Ok("save campaign", response)
	ctx.JSON(http.StatusOK, result)
}

func (controller CampaignControllerImpl) Update(ctx *gin.Context) {
	var requestID web.CampaignRequestByID
	err := ctx.ShouldBindUri(&requestID)
	if err != nil {
		controller.failedGetCampaigns(ctx, false)
		return
	}
	var request web.CampaignRequestCreate
	err = ctx.ShouldBindJSON(&request)
	if err != nil {
		errors := helper.UnprocessableEntity("failed to create campaign", err)
		ctx.JSON(http.StatusUnprocessableEntity, errors)
	}

	user := ctx.MustGet("user").(domain.User)
	request.User = user
	update, err := controller.campaignService.Update(requestID, request)
	if err != nil {
		controller.failedGetCampaigns(ctx, true)
		return
	}
	response := web.ToCampaignResponse(update)
	result := helper.Ok("update campaign", response)
	ctx.JSON(http.StatusOK, result)
}
func (controller CampaignControllerImpl) UploadImage(ctx *gin.Context) {
	var request web.CampaignImageCreate
	err := ctx.ShouldBind(&request)
	if err != nil {
		errors := helper.UnprocessableEntity("failed to create campaign image", err)
		ctx.JSON(http.StatusUnprocessableEntity, errors)
	}
	file, err := ctx.FormFile("file")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		controller.failedGetCampaignImage(ctx, false, data)
		return
	}
	user := ctx.MustGet("user").(domain.User)
	request.User = user

	path := fmt.Sprintf("images/%d-%s", user.ID, file.Filename)
	err = ctx.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		controller.failedGetCampaignImage(ctx, false, data)
		return
	}
	_, err = controller.campaignService.CreateCampaignImage(request, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		controller.failedGetCampaignImage(ctx, true, data)
		return
	}
	data := gin.H{"is_uploaded": true}
	result := helper.Ok("save campaign image", data)
	ctx.JSON(http.StatusOK, result)
}

func (controller CampaignControllerImpl) failedGetCampaigns(ctx *gin.Context, forbidden bool) {
	if forbidden {
		response := helper.Forbidden("user is not the owner of the campaign", nil)
		ctx.JSON(http.StatusForbidden, response)
	} else {
		response := helper.BadRequest("error to get campaigns", nil)
		ctx.JSON(http.StatusBadRequest, response)
	}

}

func (controller CampaignControllerImpl) failedGetCampaignImage(ctx *gin.Context, forbidden bool, data interface{}) {
	if forbidden {
		response := helper.Forbidden("user is not the owner of the campaign", data)
		ctx.JSON(http.StatusForbidden, response)
	} else {
		response := helper.BadRequest("error to upload campaign images", data)
		ctx.JSON(http.StatusBadRequest, response)
	}

}
