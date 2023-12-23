package repository

import (
	"context"
	"errors"
	"strings"
	"wanderer/features/facilities"
	"wanderer/helpers/filters"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
		if strings.Contains(queryCreate.Error.Error(), "1062") {
			return errors.New("used: facility name already exist")
		}

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

	queryUpdate := repo.mysqlDB.Where(&Facility{Id: id}).Updates(model)
	if err := queryUpdate.Error; err != nil {
		if strings.Contains(queryUpdate.Error.Error(), "1062") {
			return errors.New("used: facility name already exist")
		}

		return err
	}

	if queryUpdate.RowsAffected == 0 {
		return errors.New("not found: facility not found")
	}

	return nil
}

func (repo *facilityRepository) Delete(id uint) error {
	deleteQuery := repo.mysqlDB.Delete(&Facility{Id: id})
	if deleteQuery.Error != nil {
		if strings.Contains(deleteQuery.Error.Error(), "1451") {
			return errors.New("used: facility used by other resources")
		}

		return deleteQuery.Error
	}

	if deleteQuery.RowsAffected == 0 {
		return errors.New("not found: facility not found")
	}

	return nil
}

func (repo *facilityRepository) Import(ctx context.Context, data []facilities.Facility) error {
	var model []Facility
	for _, air := range data {
		var tmpAir = new(Facility)
		tmpAir.FromEntity(air)

		model = append(model, *tmpAir)
	}

	qry := repo.mysqlDB.WithContext(ctx).Clauses(clause.Insert{Modifier: "IGNORE"}).CreateInBatches(model, 1000)
	if qry.Error != nil {
		if strings.Contains(qry.Error.Error(), "1062") {
			return nil
		}

		return qry.Error
	}

	return nil
}
