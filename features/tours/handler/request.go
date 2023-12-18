package handler

import (
	"io"
	"time"
	"wanderer/features/facilities"
	"wanderer/features/tours"

	"github.com/labstack/echo/v4"
	"github.com/monoculum/formam/v3"
)

type TourCreateUpdateRequest struct {
	Title       string    `formam:"title"`
	Description string    `formam:"description"`
	Price       float64   `formam:"price"`
	AdminFee    float64   `formam:"admin_fee"`
	Discount    int       `formam:"discount"`
	Start       time.Time `formam:"start"`
	Finish      time.Time `formam:"finish"`
	Quota       int       `formam:"quota"`

	Thumbnail io.Reader
	Picture   []io.Reader

	Facility []uint `formam:"include_facility"`

	Itinerary []TourItineraryCreateRequest `formam:"itinerary"`

	LocationId uint `formam:"location_id"`
	AirlineId  uint `formam:"airline_id"`
}

func (req *TourCreateUpdateRequest) Bind(c echo.Context) error {
	urlVal, err := c.FormParams()
	if err != nil {
		return err
	}

	dec := formam.NewDecoder(&formam.DecoderOptions{
		TagName:           "formam",
		IgnoreUnknownKeys: true,
	})

	dec.RegisterCustomType(func(vals []string) (interface{}, error) {
		return time.Parse("2006-01-02T15:04:05Z07:00", vals[0])
	}, []interface{}{time.Time{}}, nil)

	if err := dec.Decode(urlVal, req); err != nil {
		return err
	}

	if form, err := c.MultipartForm(); err == nil {
		files := form.File["picture"]
		for _, file := range files {
			src, err := file.Open()
			if err != nil {
				return err
			}
			defer src.Close()

			req.Picture = append(req.Picture, src)
		}
	}

	thumbnail, _ := c.FormFile("thumbnail")
	if thumbnail != nil {
		src, err := thumbnail.Open()
		if err != nil {
			return err
		}
		defer src.Close()
		req.Thumbnail = src
	}

	return nil
}

func (req *TourCreateUpdateRequest) ToEntity() tours.Tour {
	var ent = new(tours.Tour)

	if req.Title != "" {
		ent.Title = req.Title
	}

	if req.Description != "" {
		ent.Description = req.Description
	}

	if req.Price != 0 {
		ent.Price = req.Price
	}

	if req.AdminFee != 0 {
		ent.AdminFee = req.AdminFee
	}

	if req.Discount != 0 {
		ent.Discount = req.Discount
	}

	if !req.Start.IsZero() {
		ent.Start = req.Start
	}

	if !req.Finish.IsZero() {
		ent.Finish = req.Finish
	}

	if req.Quota != 0 {
		ent.Quota = req.Quota
	}

	if req.Thumbnail != nil {
		ent.Thumbnail.Raw = req.Thumbnail
	}

	for _, file := range req.Picture {
		if file != nil {
			ent.Picture = append(ent.Picture, tours.File{
				Raw: file,
			})
		}
	}

	for _, facilityId := range req.Facility {
		ent.FacilityInclude = append(ent.FacilityInclude, facilities.Facility{Id: facilityId})
	}

	for _, it := range req.Itinerary {
		ent.Itinerary = append(ent.Itinerary, it.ToEntity())
	}

	if req.LocationId != 0 {
		ent.Location.Id = req.LocationId
	}

	if req.AirlineId != 0 {
		ent.Airline.Id = req.AirlineId
	}

	return *ent
}

type TourItineraryCreateRequest struct {
	Location    string `formam:"location"`
	Description string `formam:"description"`
}

func (req *TourItineraryCreateRequest) ToEntity() tours.Itinerary {
	var ent = new(tours.Itinerary)

	if req.Location != "" {
		ent.Location = req.Location
	}

	if req.Description != "" {
		ent.Description = req.Description
	}

	return *ent
}
