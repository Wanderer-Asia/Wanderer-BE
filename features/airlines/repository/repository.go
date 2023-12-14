package repository

import (
	"context"
	"errors"
	"wanderer/features/airlines"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"gorm.io/gorm"
)

func NewAirlineRepository(mysqlDB *gorm.DB, cloudinary *cloudinary.Cloudinary) airlines.Repository {
	return &airlineRepository{
		mysqlDB:    mysqlDB,
		cloudinary: cloudinary,
	}
}

type airlineRepository struct {
	mysqlDB    *gorm.DB
	cloudinary *cloudinary.Cloudinary
}

func (repo *airlineRepository) Create(newAirline airlines.Airline) error {
	if newAirline.ImageRaw != nil {
		UniqueFileName := true
		res, err := repo.cloudinary.Upload.Upload(context.TODO(), newAirline.ImageRaw, uploader.UploadParams{
			UniqueFilename: &UniqueFileName,
			Folder:         "airlines",
		})

		if err != nil {
			return err
		}

		newAirline.ImageUrl = res.URL
	}

	var model = new(Airline)
	model.FromEntity(newAirline)

	queryCreate := repo.mysqlDB.Create(model)
	if queryCreate.Error != nil {
		return queryCreate.Error
	}

	if queryCreate.RowsAffected == 0 {
		return errors.New("failed to create airline")
	}

	return nil
}

func (repo *airlineRepository) GetAll() ([]airlines.Airline, error) {
	var dataAirline []Airline
	if err := repo.mysqlDB.Find(&dataAirline).Error; err != nil {
		return nil, err
	}

	var result []airlines.Airline
	for _, airline := range dataAirline {
		result = append(result, *airline.ToEntity())
	}

	return result, nil
}

func (repo *airlineRepository) Update(id uint, updateAirline airlines.Airline) error {
	if updateAirline.ImageRaw != nil {
		UniqueFileName := true
		res, err := repo.cloudinary.Upload.Upload(context.TODO(), updateAirline.ImageRaw, uploader.UploadParams{
			UniqueFilename: &UniqueFileName,
			Folder:         "airlines",
		})

		if err != nil {
			return err
		}

		updateAirline.ImageUrl = res.URL
	}

	var model = new(Airline)
	model.FromEntity(updateAirline)

	if err := repo.mysqlDB.Where(&Airline{Id: id}).Updates(model).Error; err != nil {
		return err
	}

	return nil
}

func (repo *airlineRepository) Delete(id uint) error {
	deleteQuery := repo.mysqlDB.Delete(&Airline{Id: id})
	if deleteQuery.Error != nil {
		return deleteQuery.Error
	}

	if deleteQuery.RowsAffected == 0 {
		return errors.New("not found")
	}

	return nil
}
