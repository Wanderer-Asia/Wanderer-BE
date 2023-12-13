package repository

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id       uint   `gorm:"column:id; primaryKey;"`
	Name     string `gorm:"column:name; type:varchar(200);"`
	Phone    string `gorm:"column:phone; type:varchar(20);"`
	Email    string `gorm:"column:email; type:varchar(255); unique;"`
	Password string `gorm:"column:password; type:varchar(72); not null;"`
	Image    string `gorm:"column:image; type:text;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
