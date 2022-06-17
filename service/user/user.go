package user

import (
	"errors"
	"practice/entity"
	"practice/helper"
)

const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTMwNjQ3ODQsImlkIjo1LCJuYW1lIjoib2Rpc2EiLCJyb2xlIjoib3duZXIifQ.82JK0StJsC8iU0M3itlqxJ4w6ZBZQEGsiXt3u-iF9lc"

type UserService struct{}

func NewUserSvc() UserInterface {
	return &UserService{}
}

func (n UserService) Register(user *entity.User) (*entity.User, error) {
	if _, ok := helper.ValidMailAddress(user.Email); !ok {
		return nil, errors.New("error to validate email")
	}

	if user.Email == "" {
		return nil, errors.New("email must be filled")
	}

	if user.Username == "" {
		return nil, errors.New("username must be filled")
	}

	if user.Password == "" {
		return nil, errors.New("password must be filled ")
	}

	if len(user.Password) < 6 {
		return nil, errors.New("password must have length more than 6")
	}

	if user.Age == 0 {
		return nil, errors.New("age must be greater than nol")
	}

	return user, nil

}

//func (n UserService) Login(email, password string) (string, error) {
//	//TODO implement me
//	if _, ok := helper.ValidMailAddress(email); !ok {
//		return "", errors.New("error to validate email")
//	}
//
//	if email == "" {
//		return "", errors.New("email must be filled")
//	}
//
//	if password == "" {
//		return "", errors.New("password must be filled")
//	}
//
//	return token, nil
//}
