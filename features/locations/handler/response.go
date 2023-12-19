package handler

import (
	"time"
	"wanderer/features/locations"
	"wanderer/features/tours"
)

type LocationResponse struct {
	Id    uint   `json:"location_id,omitempty"`
	Name  string `json:"name,omitempty"`
	Image string `json:"image,omitempty"`
}

func (res *LocationResponse) FromEntity(ent locations.Location) {
	if ent.Id != 0 {
		res.Id = ent.Id
	}

	if ent.Name != "" {
		res.Name = ent.Name
	}

	if ent.ImageUrl != "" {
		res.Image = ent.ImageUrl
	} else {
		res.Image = "default"
	}
}

type TourResponse struct {
	Id        uint      `json:"tour_id,omitempty"`
	Title     string    `json:"title,omitempty"`
	Discount  int       `json:"discount"`
	Start     time.Time `json:"start,omitempty"`
	Quota     int       `json:"quota,omitempty"`
	Rating    float32   `json:"rating,omitempty"`
	Thumbnail string    `json:"thumbnail"`
	Location  string    `json:"location,omitempty"`
}

type DetailLocationResponse struct {
	Name  string `json:"name,omitempty"`
	Image string `json:"image,omitempty"`

	Tours []TourResponse `json:"tours,omitempty"`
}

func (res *DetailLocationResponse) FromEntity(ent locations.Location, tours []tours.Tour) {
	if ent.Name != "" {
		res.Name = ent.Name
	}

	if ent.ImageUrl != "" {
		res.Image = ent.ImageUrl
	} else {
		res.Image = "default"
	}

	for _, tour := range tours {
		res.Tours = append(res.Tours, TourResponse{
			Id:        tour.Id,
			Title:     tour.Title,
			Discount:  tour.Discount,
			Start:     tour.Start,
			Quota:     tour.Quota,
			Rating:    tour.Rating,
			Thumbnail: tour.Thumbnail.Url,
			Location:  ent.Name,
		})
	}
}
