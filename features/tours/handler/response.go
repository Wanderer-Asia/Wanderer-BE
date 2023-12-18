package handler

import (
	"time"
	"wanderer/features/tours"
)

type TourResponse struct {
	Id          uint      `json:"tour_id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Price       float64   `json:"price"`
	AdminFee    float64   `json:"admin_fee"`
	Discount    int       `json:"discount"`
	Start       time.Time `json:"start,omitempty"`
	Finish      time.Time `json:"finish,omitempty"`
	Quota       int       `json:"quota,omitempty"`
	Available   int       `json:"available,omitempty"`
	Rating      float32   `json:"rating,omitempty"`

	Thumbnail string   `json:"thumbnail"`
	Picture   []string `json:"picture"`

	Facility struct {
		Include []string `json:"include"`
		Exclude []string `json:"exclude"`
	} `json:"facility"`

	Itinerary []TourItineraryResponse `json:"itinerary"`

	Location TourLocationResponse `json:"location"`
	Airline  TourAirlineResponse  `json:"airline"`

	Reviews []TourReviewResponse `json:"reviews"`
}

func (res *TourResponse) FromEntity(ent tours.Tour) {
	if ent.Id != 0 {
		res.Id = ent.Id
	}

	res.Title = ent.Title
	res.Description = ent.Description
	res.Price = ent.Price
	res.AdminFee = ent.AdminFee
	res.Discount = ent.Discount
	res.Start = ent.Start
	res.Finish = ent.Finish
	res.Quota = ent.Quota
	res.Available = ent.Available
	res.Rating = ent.Rating

	if ent.Thumbnail.Url != "" {
		res.Thumbnail = ent.Thumbnail.Url
	} else {
		res.Thumbnail = "default"
	}

	for _, pict := range ent.Picture {
		if pict.Url != "" {
			res.Picture = append(res.Picture, pict.Url)
		}
	}

	for _, fac := range ent.FacilityInclude {
		res.Facility.Include = append(res.Facility.Include, fac.Name)
	}

	for _, fac := range ent.FacilityExclude {
		res.Facility.Exclude = append(res.Facility.Exclude, fac.Name)
	}

	for _, it := range ent.Itinerary {
		var tmpItinerary = new(TourItineraryResponse)
		tmpItinerary.FromEntity(it)

		res.Itinerary = append(res.Itinerary, *tmpItinerary)
	}

	res.Location = TourLocationResponse{Name: ent.Location.Name}
	res.Airline = TourAirlineResponse{Name: ent.Airline.Name}
}

type TourItineraryResponse struct {
	Location    string `json:"location"`
	Description string `json:"description"`
}

func (res *TourItineraryResponse) FromEntity(ent tours.Itinerary) {
	res.Location = ent.Location
	res.Description = ent.Description
}

type TourLocationResponse struct {
	Name string `json:"name"`
}

type TourAirlineResponse struct {
	Name string `json:"name"`
}

type TourReviewResponse struct {
	User      TourUserResponse `json:"user"`
	Text      string           `json:"text,omitempty"`
	CreatedAt time.Time        `json:"created_at"`
}

type TourUserResponse struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}
