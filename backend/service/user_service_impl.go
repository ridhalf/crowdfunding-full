package service

import (
	"crowdfunding/helper"
	"crowdfunding/model/domain"
	"crowdfunding/model/web"
	"crowdfunding/repository"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserServiceImpl(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
	}
}

func (service UserServiceImpl) Register(request web.UserRequestRegister) (domain.User, error) {
	user := domain.User{}
	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)

	if err != nil {
		return user, err
	}

	user = domain.User{
		Name:         request.Name,
		Email:        request.Email,
		Occupation:   request.Occupation,
		Role:         "user",
		PasswordHash: string(password),
	}

	save, err := service.UserRepository.Save(user)

	return helper.ResultOrError(save, err)
}

func (service UserServiceImpl) Login(request web.UserRequestLogin) (domain.User, error) {
	//TODO implement me
	email := request.Email
	password := request.Password

	user, err := service.UserRepository.FindByEmail(email)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("the email or password you entered is incorrect")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, errors.New("the email or password you entered is incorrect")
	}

	return user, nil
}

func (service UserServiceImpl) IsEmailAvailable(request web.UserRequestEmailCheck) (bool, error) {
	email := request.Email
	user, err := service.UserRepository.FindByEmail(email)
	if err != nil {
		return false, err
	}
	if user.ID == 0 {
		return true, nil
	}
	return false, nil
}

func (service UserServiceImpl) SaveAvatar(ID int, fileLocation string) (domain.User, error) {
	id := ID
	user, err := service.UserRepository.FindByID(id)
	if err != nil {
		return user, err
	}
	user.AvatarFileName = fileLocation

	result, err := service.UserRepository.Update(user)
	return helper.ResultOrError(result, err)

}

func (service UserServiceImpl) FindById(ID int) (domain.User, error) {
	user := domain.User{}
	user, err := service.UserRepository.FindByID(ID)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("the id of the user is not found")
	}
	return user, nil
}
