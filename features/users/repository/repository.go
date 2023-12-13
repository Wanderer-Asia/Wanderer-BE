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

func (repo *userRepository) Register(newUser users.User) error {
	var model = new(User)
	model.FromEntity(newUser)

	if err := repo.mysqlDB.Create(model).Error; err != nil {
		return err
	}

	return nil
}

func (repo *userRepository) Login(email string) (*users.User, error) {
	panic("unimplemented")
}

func (repo *userRepository) Update(id uint, updateUser users.User) error {
	panic("unimplemented")
}

func (repo *userRepository) Delete(id uint) error {
	panic("unimplemented")
}
