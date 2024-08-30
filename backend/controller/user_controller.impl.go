package controller

import (
	"crowdfunding/auth"
	"crowdfunding/helper"
	"crowdfunding/model/domain"
	"crowdfunding/model/web"
	"crowdfunding/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserControllerImpl struct {
	userService service.UserService
	authJwt     auth.JwtService
}

func NewUserController(userService service.UserService, authJwt auth.JwtService) UserController {
	return &UserControllerImpl{
		userService: userService,
		authJwt:     authJwt,
	}
}

func (controller *UserControllerImpl) Register(ctx *gin.Context) {
	registerRequest := web.UserRequestRegister{}
	err := ctx.ShouldBindJSON(&registerRequest)
	if err != nil {
		response := helper.UnprocessableEntity("register account failed", err)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	user, err := controller.userService.Register(registerRequest)
	if err != nil {
		controller.registerFailedResponse(ctx)
		return
	}

	token, err := controller.authJwt.GenerateToken(user.ID)
	if err != nil {
		controller.registerFailedResponse(ctx)
		return
	}

	response := web.ToUserResponse(user, token)
	result := helper.Ok("Account has been registered", response)
	ctx.JSON(http.StatusOK, result)
}

func (controller *UserControllerImpl) Login(ctx *gin.Context) {
	//TODO implement me
	loginRequest := web.UserRequestLogin{}
	err := ctx.ShouldBindJSON(&loginRequest)
	if err != nil {
		response := helper.UnprocessableEntity("login account failed", err)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	login, err := controller.userService.Login(loginRequest)
	if err != nil {
		response := helper.UnprocessableEntityString("login account failed", err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := controller.authJwt.GenerateToken(login.ID)
	if err != nil {
		response := helper.BadRequest("login account failed", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := web.ToUserResponse(login, token)
	result := helper.Ok(`login successful. welcome back!`, response)
	ctx.JSON(http.StatusOK, result)
}

func (controller *UserControllerImpl) IsEmailAvailable(ctx *gin.Context) {

	emailCheck := web.UserRequestEmailCheck{}
	err := ctx.ShouldBindJSON(&emailCheck)
	if err != nil {
		response := helper.UnprocessableEntity("email check failed", err)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	available, err := controller.userService.IsEmailAvailable(emailCheck)
	if err != nil {
		response := helper.UnprocessableEntityString("an unexpected error occurred. Please try again later", err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	data := gin.H{"is_available": available}
	var metaString string
	if available {
		metaString = "email is available"
	} else {
		metaString = "email is not available"
	}
	result := helper.Ok(metaString, data)
	ctx.JSON(http.StatusOK, result)
}

func (controller *UserControllerImpl) UploadAvatar(ctx *gin.Context) {
	file, err := ctx.FormFile("avatar")
	if err != nil {
		controller.uploadAvatarFailedResponse(ctx)
		return
	}

	//=====DAPAT DARI JWT
	user := ctx.MustGet("user").(domain.User)
	userID := user.ID
	//=====

	path := "images/" + strconv.Itoa(userID) + "-" + file.Filename
	err = ctx.SaveUploadedFile(file, path)
	if err != nil {
		controller.uploadAvatarFailedResponse(ctx)
		return
	}

	_, err = controller.userService.SaveAvatar(userID, path)
	if err != nil {
		controller.uploadAvatarFailedResponse(ctx)
		return
	}

	data := gin.H{"is_uploaded": true}
	result := helper.Ok("avatar uploaded successfully", data)
	ctx.JSON(http.StatusOK, result)
	return
}

func (controller *UserControllerImpl) FetchUser(ctx *gin.Context) {
	user := ctx.MustGet("user").(domain.User)
	response := web.ToUserResponse(user, "")
	result := helper.Ok("avatar uploaded successfully", response)
	ctx.JSON(http.StatusOK, result)
}

func (controller *UserControllerImpl) uploadAvatarFailedResponse(ctx *gin.Context) {
	data := gin.H{"is_uploaded": false}
	response := helper.BadRequest("upload avatar failed", data)
	ctx.JSON(http.StatusBadRequest, response)
}
func (controller *UserControllerImpl) registerFailedResponse(ctx *gin.Context) {
	response := helper.BadRequest("register account failed", nil)
	ctx.JSON(http.StatusBadRequest, response)
}
