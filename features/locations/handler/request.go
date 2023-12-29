package handler

import (
	"encoding/csv"
	"io"
	"strings"
	"wanderer/features/locations"

	echo "github.com/labstack/echo/v4"
)

type LocationCreateUpdateRequest struct {
	Name     string `form:"name"`
	ImageRaw io.Reader
}

func (req *LocationCreateUpdateRequest) ToEntity() locations.Location {
	var ent = new(locations.Location)

	if req.Name != "" {
		ent.Name = req.Name
	}

	if req.ImageRaw != nil {
		ent.ImageRaw = req.ImageRaw
	}

	return *ent
}

type ImportLocationRequest struct {
	File io.Reader
}

func (req *ImportLocationRequest) Bind(c echo.Context) error {
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

func (req *ImportLocationRequest) ToEntity() ([]locations.Location, error) {
	var ent []locations.Location

	if req.File != nil {
		reader := csv.NewReader(req.File)
		records, err := reader.ReadAll()
		if err != nil {
			return nil, err
		}

		for i := 1; i < len(records); i++ {
			var tmpLocation = new(locations.Location)

			value := strings.Split(records[i][0], ";")
			if len(value) >= 1 && value[0] != "" {
				tmpLocation.Name = value[0]
			} else {
				continue
			}

			if len(value) >= 2 && value[1] != "" {
				tmpLocation.ImageUrl = value[1]
			}

			ent = append(ent, *tmpLocation)
		}
	}

	return ent, nil
}
