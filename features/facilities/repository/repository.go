package repository

import (
	"errors"
	"wanderer/features/facilities"
	"wanderer/helpers/filters"

	"gorm.io/gorm"
)

func NewFacilityRepository(mysqlDB *gorm.DB) facilities.Repository {
	return &facilityRepository{
		mysqlDB: mysqlDB,
	}
}

type facilityRepository struct {
	mysqlDB *gorm.DB
}

func (repo *facilityRepository) Create(newFacility facilities.Facility) error {
	var model = new(Facility)
	model.FromEntity(newFacility)

	queryCreate := repo.mysqlDB.Create(model)
	if queryCreate.Error != nil {
		return queryCreate.Error
	}

	if queryCreate.RowsAffected == 0 {
		return errors.New("failed to create facility")
	}

	return nil
}

func (repo *facilityRepository) GetAll(flt filters.Filter) ([]facilities.Facility, error) {
	panic("unimplemented")
}

func (repo *facilityRepository) Update(id uint, updateFacility facilities.Facility) error {
	panic("unimplemented")
}

func (repo *facilityRepository) Delete(id uint) error {
	panic("unimplemented")
}
