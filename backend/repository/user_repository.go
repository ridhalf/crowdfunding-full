package repository

import (
	"crowdfunding/model/domain"
)

type UserRepository interface {
	Save(user domain.User) (domain.User, error)
	FindByEmail(email string) (domain.User, error)
	FindByID(ID int) (domain.User, error)
	Update(user domain.User) (domain.User, error)
}
