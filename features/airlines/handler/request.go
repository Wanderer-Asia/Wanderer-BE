package handler

import (
	"encoding/csv"
	"io"
	"strings"
	"wanderer/features/airlines"

	echo "github.com/labstack/echo/v4"
)

type CreateRequest struct {
	Name  string `form:"name"`
	Image io.Reader
}

func (req *CreateRequest) ToEntity() *airlines.Airline {
	var ent = new(airlines.Airline)

	if req.Name != "" {
		ent.Name = req.Name
	}

	if req.Image != nil {
		ent.ImageRaw = req.Image
	}

	return ent
}

type ImportAirlineRequest struct {
	File io.Reader
}

func (req *ImportAirlineRequest) Bind(c echo.Context) error {
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

func (req *ImportAirlineRequest) ToEntity() ([]airlines.Airline, error) {
	var ent []airlines.Airline

	if req.File != nil {
		reader := csv.NewReader(req.File)
		records, err := reader.ReadAll()
		if err != nil {
			return nil, err
		}

		for i := 1; i < len(records); i++ {
			var tmpAirline = new(airlines.Airline)

			value := strings.Split(records[i][0], ";")
			if len(value) >= 1 && value[0] != "" {
				tmpAirline.Name = value[0]
			} else {
				continue
			}

			if len(value) >= 2 && value[1] != "" {
				tmpAirline.ImageUrl = value[1]
			}

			ent = append(ent, *tmpAirline)
		}
	}

	return ent, nil
}
