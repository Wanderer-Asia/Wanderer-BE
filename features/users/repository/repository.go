package repository

import (
	"wanderer/features/users"

	"github.com/cloudinary/cloudinary-go/v2"
	"gorm.io/gorm"
)

func NewUserRepository(mysqlDB *gorm.DB, cloudinary *cloudinary.Cloudinary) users.Repository {
	return &userRepository{
		mysqlDB:    mysqlDB,
		cloudinary: cloudinary,
	}
}

type userRepository struct {
	mysqlDB    *gorm.DB
	cloudinary *cloudinary.Cloudinary
}

func (repo *userRepository) Register() error {
	panic("unimplemented")
}

func (repo *userRepository) Login() (*users.User, error) {
	panic("unimplemented")
}

func (repo *userRepository) Update() error {
	panic("unimplemented")
}

func (repo *userRepository) Delete() error {
	panic("unimplemented")
}
