package repository

import (
	"context"
	"errors"
	"strings"
	"wanderer/features/airlines"
	"wanderer/helpers/filters"
	"wanderer/utils/files"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func NewAirlineRepository(mysqlDB *gorm.DB, cloud files.Cloud) airlines.Repository {
	return &airlineRepository{
		mysqlDB: mysqlDB,
		cloud:   cloud,
	}
}

type airlineRepository struct {
	mysqlDB *gorm.DB
	cloud   files.Cloud
}

func (repo *airlineRepository) Create(newAirline airlines.Airline) error {
	if newAirline.ImageRaw != nil {
		url, err := repo.cloud.Upload(context.Background(), "airlines", newAirline.ImageRaw)
		if err != nil {
			return err
		}

		newAirline.ImageUrl = *url
	}

	var model = new(Airline)
	model.FromEntity(newAirline)

	queryCreate := repo.mysqlDB.Create(model)
	if queryCreate.Error != nil {
		if strings.Contains(queryCreate.Error.Error(), "1062") {
			return errors.New("used: airline name already exist")
		}

		return queryCreate.Error
	}

	if queryCreate.RowsAffected == 0 {
		return errors.New("failed to create airline")
	}

	return nil
}

func (repo *airlineRepository) GetAll(flt filters.Filter) ([]airlines.Airline, error) {
	var dataAirline []Airline
	qry := repo.mysqlDB

	if flt.Search.Keyword != "" {
		qry = qry.Where("name like ?", "%"+flt.Search.Keyword+"%")
	}

	qry = qry.Find(&dataAirline)
	if qry.Error != nil {
		return nil, qry.Error
	}

	var result []airlines.Airline
	for _, airline := range dataAirline {
		result = append(result, *airline.ToEntity())
	}

	return result, nil
}

func (repo *airlineRepository) Update(id uint, updateAirline airlines.Airline) error {
	if updateAirline.ImageRaw != nil {
		url, err := repo.cloud.Upload(context.Background(), "airlines", updateAirline.ImageRaw)
		if err != nil {
			return err
		}

		updateAirline.ImageUrl = *url
	}

	var model = new(Airline)
	model.FromEntity(updateAirline)

	updateQuery := repo.mysqlDB.Where(&Airline{Id: id}).Updates(model)
	if err := updateQuery.Error; err != nil {
		if strings.Contains(err.Error(), "1062") {
			return errors.New("used: airline name already exist")
		}

		return err
	}

	if updateQuery.RowsAffected == 0 {
		return errors.New("not found: airline not found")
	}

	return nil
}

func (repo *airlineRepository) Delete(id uint) error {
	deleteQuery := repo.mysqlDB.Delete(&Airline{Id: id})
	if deleteQuery.Error != nil {
		if strings.Contains(deleteQuery.Error.Error(), "1451") {
			return errors.New("used: airline used by other resources")
		}

		return deleteQuery.Error
	}

	if deleteQuery.RowsAffected == 0 {
		return errors.New("not found: airline not found")
	}

	return nil
}

func (repo *airlineRepository) Import(ctx context.Context, data []airlines.Airline) error {
	var model []Airline
	for _, air := range data {
		var tmpAir = new(Airline)
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
