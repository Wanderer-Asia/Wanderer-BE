package handler

import "wanderer/features/users"

type RegisterRequest struct {
	Name     string `json:"name,omitempty"`
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
