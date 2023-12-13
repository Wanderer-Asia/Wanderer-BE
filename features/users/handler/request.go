package handler

type UserResponse struct {
	Id    uint   `json:"user_id,omitempty"`
	Name  string `json:"name,omitempty"`
	Phone string `json:"phone,omitempty"`
	Email string `json:"email,omitempty"`
	Image string `json:"image,omitempty"`
	Role  string `json:"role,omitempty"`
	Token string `json:"token,omitempty"`
}
