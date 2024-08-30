package service

import (
	"crowdfunding/model/domain"
	"crowdfunding/model/web"
)

type UserService interface {
	Register(request web.UserRequestRegister) (domain.User, error)
	Login(request web.UserRequestLogin) (domain.User, error)
	IsEmailAvailable(request web.UserRequestEmailCheck) (bool, error)
	SaveAvatar(ID int, fileLocation string) (domain.User, error)
	FindById(ID int) (domain.User, error)
}
