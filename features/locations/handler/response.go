package handler

import (
	"time"
	"wanderer/features/locations"
)

type LocationResponse struct {
	Id    uint   `json:"location_id,omitempty"`
	Name  string `json:"name,omitempty"`
	Image string `json:"image,omitempty"`

	Tours []TourResponse `json:"tours,omitempty"`
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

	if len(ent.Tours) != 0 {
		for _, tour := range ent.Tours {
			var tmpTour = new(TourResponse)
			tmpTour.FromEntity(tour, ent)

			res.Tours = append(res.Tours, *tmpTour)
		}
	}
}

type TourResponse struct {
	Id        uint             `json:"tour_id"`
	Title     string           `json:"title,omitempty"`
	Discount  int              `json:"discount,omitempty"`
	Start     time.Time        `json:"start,omitempty"`
	Quota     int              `json:"quota"`
	Rating    float32          `json:"rating,omitempty"`
	Thumbnail string           `json:"thumbnail"`
	Location  LocationResponse `json:"location,omitempty"`
}

func (res *TourResponse) FromEntity(ent locations.Tour, loc locations.Location) {
	if ent.Id != 0 {
		res.Id = ent.Id
	}

	if ent.Title != "" {
		res.Title = ent.Title
	}

	if ent.Discount != 0 {
		res.Discount = ent.Discount
	}

	if !ent.Start.IsZero() {
		res.Start = ent.Start
	}

	if ent.Quota != 0 {
		res.Quota = ent.Quota
	}

	if ent.Rating != 0 {
		res.Rating = ent.Rating
	}

	if ent.Thumbnail != "" {
		res.Thumbnail = ent.Thumbnail
	} else {
		res.Thumbnail = "https://res.cloudinary.com/dhxzinjxp/image/upload/v1703490571/asset-default/tour_zk9i73.png"
	}

	if loc.Name != "" {
		res.Location.Name = loc.Name
	}
}
