package repository

import (
	"context"
	"errors"
	"wanderer/features/users"
	"wanderer/utils/files"

	"gorm.io/gorm"
)

func NewUserRepository(mysqlDB *gorm.DB, cloud files.Cloud) users.Repository {
	return &userRepository{
		mysqlDB: mysqlDB,
		cloud:   cloud,
	}
}

type userRepository struct {
	mysqlDB *gorm.DB
	cloud   files.Cloud
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
		url, err := repo.cloud.Upload(context.Background(), "users", updateUser.ImageRaw)
		if err != nil {
			return err
		}

		updateUser.ImageUrl = *url
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

func (repo *userRepository) GetById(id uint) (*users.User, error) {
	var model = new(User)

	if err := repo.mysqlDB.Preload("Bookings").Preload("Bookings.BookingDetails").Preload("Bookings.Tour").Preload("Reviews").Where("id = ?", id).First(model).Error; err != nil {
		return nil, err
	}

	return model.ToEntity(), nil
}
