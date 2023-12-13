package users

import "io"

type User struct {
	Id       uint
	Name     string
	Email    string
	Phone    string
	Password string
	Role     string

	ImageUrl string
	ImageRaw io.Reader
}

type Handler interface{}

type Service interface{}

type Repository interface{}
