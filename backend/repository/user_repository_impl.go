package repository

import (
	"crowdfunding/helper"
	"crowdfunding/model/domain"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db}
}

func (repository UserRepositoryImpl) Save(user domain.User) (domain.User, error) {
	//TODO implement me
	err := repository.db.Create(&user).Error
	return helper.ResultOrError(user, err)
}

func (repository UserRepositoryImpl) FindByEmail(email string) (domain.User, error) {
	user := domain.User{}
	err := repository.db.Where("email = ?", email).Find(&user).Error
	return helper.ResultOrError(user, err)
}

func (repository UserRepositoryImpl) FindByID(ID int) (domain.User, error) {
	//TODO implement me
	user := domain.User{}
	err := repository.db.Where("id = ?", ID).Find(&user).Error
	return helper.ResultOrError(user, err)
}

func (repository UserRepositoryImpl) Update(user domain.User) (domain.User, error) {
	err := repository.db.Save(&user).Error
	return helper.ResultOrError(user, err)
}
