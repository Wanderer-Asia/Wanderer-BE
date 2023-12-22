package handler

import (
	"io"
	"wanderer/features/users"
)

type RegisterRequest struct {
	Name     string `json:"fullname,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type LoginRequest struct {
	Name     string `json:"fullname,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func (req *RegisterRequest) ToEntity() *users.User {
	var ent = new(users.User)

	if req.Name != "" {
		ent.Name = req.Name
	}

	if req.Phone != "" {
		ent.Phone = req.Phone
	}

	if req.Email != "" {
		ent.Email = req.Email
	}

	if req.Password != "" {
		ent.Password = req.Password
	}

	return ent
}

func (req *LoginRequest) ToEntity() *users.User {
	var ent = new(users.User)

	if req.Email != "" {
		ent.Email = req.Email
	}

	if req.Password != "" {
		ent.Password = req.Password
	}

	return ent
}

type UpdateRequest struct {
	Name     string `form:"fullname"`
	Phone    string `form:"phone"`
	Email    string `form:"email"`
	Password string `form:"password"`
	Image    io.Reader
}

func (req *UpdateRequest) ToEntity() *users.User {
	var ent = new(users.User)

	if req.Name != "" {
		ent.Name = req.Name
	}

	if req.Phone != "" {
		ent.Phone = req.Phone
	}

	if req.Email != "" {
		ent.Email = req.Email
	}

	if req.Password != "" {
		ent.Password = req.Password
	}

	if req.Image != nil {
		ent.ImageRaw = req.Image
	}

	return ent
}
