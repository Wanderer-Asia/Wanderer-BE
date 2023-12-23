package handler

import (
	"encoding/csv"
	"io"
	"strings"
	"wanderer/features/facilities"

	"github.com/labstack/echo/v4"
)

type CreateRequest struct {
	Name string `form:"name"`
}

func (req *CreateRequest) ToEntity() *facilities.Facility {
	var ent = new(facilities.Facility)

	if req.Name != "" {
		ent.Name = req.Name
	}

	return ent
}

type ImportFacilityRequest struct {
	File io.Reader
}

func (req *ImportFacilityRequest) Bind(c echo.Context) error {
	File, err := c.FormFile("file")
	if err != nil {
		return err
	}

	src, err := File.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	req.File = src

	return nil
}

func (req *ImportFacilityRequest) ToEntity() ([]facilities.Facility, error) {
	var ent []facilities.Facility

	if req.File != nil {
		reader := csv.NewReader(req.File)
		records, err := reader.ReadAll()
		if err != nil {
			return nil, err
		}

		for i := 1; i < len(records); i++ {
			var tmpFacility = new(facilities.Facility)

			value := strings.Split(records[i][0], ";")
			if len(value) >= 1 && value[0] != "" {
				tmpFacility.Name = value[0]
			} else {
				continue
			}

			ent = append(ent, *tmpFacility)
		}
	}

	return ent, nil
}
