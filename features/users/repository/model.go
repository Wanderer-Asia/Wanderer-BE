package repository

import (
	"time"
	"wanderer/features/users"

	"gorm.io/gorm"
)

type User struct {
	Id       uint   `gorm:"column:id; primaryKey;"`
	Name     string `gorm:"column:fullname; type:varchar(200);"`
	Phone    string `gorm:"column:phone; type:varchar(20);"`
	Email    string `gorm:"column:email; type:varchar(255); unique;"`
	Password string `gorm:"column:password; type:varchar(72); not null;"`
	Image    string `gorm:"column:image; type:text;"`
	Role     string `gorm:"column:role; type:enum('admin', 'user');"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (mod *User) FromEntity(ent users.User) {
	if ent.Name != "" {
		mod.Name = ent.Name
	}

	if ent.Phone != "" {
		mod.Phone = ent.Phone
	}

	if ent.Email != "" {
		mod.Email = ent.Email
	}

	if ent.Password != "" {
		mod.Password = ent.Password
	}

	if ent.ImageUrl != "" {
		mod.Image = ent.ImageUrl
	}

	if ent.Role != "" {
		mod.Role = ent.Role
	}
}

func (mod *User) ToEntity() *users.User {
	var ent = new(users.User)

	if mod.Id != 0 {
		ent.Id = mod.Id
	}

	if mod.Id != 0 {
		ent.Id = mod.Id
	}

	if mod.Name != "" {
		ent.Name = mod.Name
	}

	if mod.Phone != "" {
		ent.Phone = mod.Phone
	}

	if mod.Email != "" {
		ent.Email = mod.Email
	}

	if mod.Password != "" {
		ent.Password = mod.Password
	}

	if mod.Image != "" {
		ent.ImageUrl = mod.Image
	}

	if mod.Role != "" {
		ent.Role = mod.Role
	}

	if !mod.CreatedAt.IsZero() {
		ent.CreatedAt = mod.CreatedAt
	}

	if !mod.UpdatedAt.IsZero() {
		ent.UpdatedAt = mod.UpdatedAt
	}

	if !mod.DeletedAt.Time.IsZero() {
		ent.DeletedAt = mod.DeletedAt.Time
	}

	return ent
}
