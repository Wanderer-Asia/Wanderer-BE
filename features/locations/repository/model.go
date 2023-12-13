package repository

import (
	"time"

	"gorm.io/gorm"
)

type Location struct {
	Id   uint   `gorm:"column:id; primaryKey;"`
	Name string `gorm:"column:name; type:varchar(200);"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
