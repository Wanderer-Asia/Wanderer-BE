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

	if err := srv.repo.Register(newUser); err != nil {
		return err
	}

	return nil
}

func (srv *userService) Login(email string, password string) (*users.User, error) {
	if email == "" {
		return nil, errors.New("validate: email can't be empty")
	}

	if password == "" {
		return nil, errors.New("validate: password can't be empty")
	}

	result, err := srv.repo.Login(email)
	if err != nil {
		return nil, err
	}

	if err := srv.enc.Compare(result.Password, password); err != nil {
		return nil, errors.New("validate: wrong password")
	}

	return result, nil
}

func (srv *userService) Update(id uint, updateUser users.User) error {
	if id == 0 {
		return errors.New("validate: invalid user id")
	}

	if updateUser.Password != "" {
		hash, err := srv.enc.Hash(updateUser.Password)
		if err != nil {
			return err
		}

		updateUser.Password = hash
	}

	if err := srv.repo.Update(id, updateUser); err != nil {
		return err
	}

	return nil
}

func (srv *userService) Delete(id uint) error {
	if id == 0 {
		return errors.New("validate: invalid user id")
	}

	if err := srv.repo.Delete(id); err != nil {
		return err
	}

	return nil
}

func (srv *userService) Detail(id uint) (*users.User, error) {
	if id == 0 {
		return nil, errors.New("validate: invalid user id")
	}

	result, err := srv.repo.Detail(id)
	if err != nil {
		return nil, err
	}

	return result, nil
}
