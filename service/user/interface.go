package user

import "practice/entity"

type UserInterface interface {
	Register(user *entity.User) (*entity.User, error)
	//Login(email, password string) (string, error)
}
