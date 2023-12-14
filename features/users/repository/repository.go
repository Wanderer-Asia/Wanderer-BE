package repository

import (
	"context"
	"errors"
	"wanderer/features/users"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
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
	var model = new(User)

	if err := repo.mysqlDB.Where("email = ?", email).First(model).Error; err != nil {
		return nil, err
	}

	return model.ToEntity(), nil
}

func (repo *userRepository) Update(id uint, updateUser users.User) error {
	if updateUser.ImageRaw != nil {
		UniqueFileName := true
		res, err := repo.cloudinary.Upload.Upload(context.TODO(), updateUser.ImageRaw, uploader.UploadParams{
			UniqueFilename: &UniqueFileName,
			Folder:         "users",
		})

		if err != nil {
			return err
		}

		updateUser.ImageUrl = res.URL
	}

	var model = new(User)
	model.FromEntity(updateUser)

	if err := repo.mysqlDB.Where(&User{Id: id}).Updates(model).Error; err != nil {
		return err
	}

	return nil
}

func (repo *userRepository) Delete(id uint) error {
	deleteQuery := repo.mysqlDB.Delete(&User{Id: id})
	if deleteQuery.Error != nil {
		return deleteQuery.Error
	}

	if deleteQuery.RowsAffected == 0 {
		return errors.New("not found")
	}

	return nil
}
