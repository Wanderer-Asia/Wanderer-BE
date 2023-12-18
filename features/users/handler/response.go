package handler

import "wanderer/features/users"

type UserResponse struct {
	Id    uint   `json:"user_id,omitempty"`
	Name  string `json:"fullname,omitempty"`
	Phone string `json:"phone,omitempty"`
	Email string `json:"email,omitempty"`
	Image string `json:"image,omitempty"`
	Role  string `json:"role,omitempty"`
	Token string `json:"token,omitempty"`
}

type LoginResponse struct {
	Id    uint   `json:"user_id,omitempty"`
	Name  string `json:"fullname,omitempty"`
	Image string `json:"image,omitempty"`
	Role  string `json:"role,omitempty"`
	Token string `json:"token,omitempty"`
}

func (res *LoginResponse) FromEntity(ent users.User) {
	if ent.Id != 0 {
		res.Id = ent.Id
	}

	if ent.Name != "" {
		res.Name = ent.Name
	}

	if ent.ImageUrl != "" {
		res.Image = ent.ImageUrl
	} else {
		res.Image = "default"
	}

	if ent.Role != "" {
		res.Role = ent.Role
	}

}
