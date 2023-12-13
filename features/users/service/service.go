package service

import (
	"wanderer/features/users"
	"wanderer/helpers/encrypt"
)

func NewUserService(repo users.Repository, enc encrypt.BcryptHash) users.Service {
	return &userService{
		repo: repo,
		enc:  enc,
	}
}

type userService struct {
	repo users.Repository
	enc  encrypt.BcryptHash
}

func (srv *userService) Register() error {
	panic("unimplemented")
}

func (srv *userService) Login() (*users.User, error) {
	panic("unimplemented")
}

func (srv *userService) Update() error {
	panic("unimplemented")
}

func (srv *userService) Delete() error {
	panic("unimplemented")
}
