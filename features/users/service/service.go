package service

import (
	"errors"
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

func (srv *userService) Register(newUser users.User) error {
	if newUser.Name == "" {
		return errors.New("validate: name can't be empty")
	}
	if newUser.Phone == "" {
		return errors.New("validate: phone can't be empty")
	}
	if newUser.Email == "" {
		return errors.New("validate: email can't be empty")
	}
	if newUser.Password == "" {
		return errors.New("validate: password can't be empty")
	}

	encrypt, err := srv.enc.Hash(newUser.Password)
	if err != nil {
		return err
	}

	newUser.Password = encrypt
	newUser.Role = "user"
	newUser.ImageUrl = "default"

	if err := srv.repo.Register(newUser); err != nil {
		return err
	}

	return nil
}

func (srv *userService) Login(email string, password string) (*users.User, error) {
	panic("unimplemented")
}

func (srv *userService) Update(id uint, updateUser users.User) error {
	panic("unimplemented")
}

func (srv *userService) Delete(id uint) error {
	panic("unimplemented")
}
