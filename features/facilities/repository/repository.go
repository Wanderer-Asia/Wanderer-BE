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
	var dataFacilities []Facility
	qry := repo.mysqlDB

	if flt.Search.Keyword != "" {
		qry = qry.Where("name like ?", "%"+flt.Search.Keyword+"%")
	}

	qry = qry.Find(&dataFacilities)
	if qry.Error != nil {
		return nil, qry.Error
	}

	var result []facilities.Facility
	for _, facility := range dataFacilities {
		result = append(result, *facility.ToEntity())
	}

	return result, nil
}

func (repo *facilityRepository) Update(id uint, updateFacility facilities.Facility) error {
	var model = new(Facility)
	model.FromEntity(updateFacility)

	if err := repo.mysqlDB.Where(&Facility{Id: id}).Updates(model).Error; err != nil {
		return err
	}

	return nil
}

func (repo *facilityRepository) Delete(id uint) error {
	deleteQuery := repo.mysqlDB.Delete(&Facility{Id: id})
	if deleteQuery.Error != nil {
		return deleteQuery.Error
	}

	if deleteQuery.RowsAffected == 0 {
		return errors.New("not found")
	}

	return nil
}
