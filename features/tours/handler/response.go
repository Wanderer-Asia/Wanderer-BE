package handler

import (
	"reflect"
	"time"
	"wanderer/features/tours"
)

type TourResponse struct {
	Id          uint       `json:"tour_id,omitempty"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	Price       float64    `json:"price"`
	AdminFee    *float64   `json:"admin_fee,omitempty"`
	Discount    float64    `json:"discount"`
	Start       time.Time  `json:"start,omitempty"`
	Finish      *time.Time `json:"finish,omitempty"`
	Quota       int        `json:"quota,omitempty"`
	Available   int        `json:"available,omitempty"`
	Rating      float32    `json:"rating"`

	Thumbnail string   `json:"thumbnail"`
	Picture   []string `json:"picture,omitempty"`

	Facility *struct {
		Include   []string `json:"include"`
		IncludeId []uint   `json:"include_id"`
		Exclude   []string `json:"exclude"`
	} `json:"facility"`

	Itinerary []ItineraryResponse `json:"itinerary,omitempty"`

	Location LocationResponse `json:"location"`
	Airline  *AirlineResponse `json:"airline,omitempty"`

	Reviews []ReviewResponse `json:"reviews,omitempty"`
}

func (res *TourResponse) FromEntity(ent tours.Tour, discountCurrency bool) {
	if ent.Id != 0 {
		res.Id = ent.Id
	}

	res.Title = ent.Title
	res.Description = ent.Description
	res.Price = ent.Price
	if ent.AdminFee != -1 {
		res.AdminFee = &ent.AdminFee
	}
	if discountCurrency {
		res.Discount = float64(ent.Discount) / 100 * ent.Price
	} else {
		res.Discount = float64(ent.Discount)
	}
	res.Start = ent.Start
	if !ent.Finish.IsZero() {
		res.Finish = &ent.Finish
	}

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

	if len(ent.FacilityInclude) != 0 || len(ent.FacilityExclude) != 0 {
		res.Facility = &struct {
			Include   []string `json:"include"`
			IncludeId []uint   `json:"include_id"`
			Exclude   []string `json:"exclude"`
		}{}
	}

	for _, fac := range ent.FacilityInclude {
		res.Facility.IncludeId = append(res.Facility.IncludeId, fac.Id)
		res.Facility.Include = append(res.Facility.Include, fac.Name)
	}

	for _, fac := range ent.FacilityExclude {
		res.Facility.Exclude = append(res.Facility.Exclude, fac.Name)
	}

	for _, it := range ent.Itinerary {
		var tmpItinerary = new(ItineraryResponse)
		tmpItinerary.FromEntity(it)

		res.Itinerary = append(res.Itinerary, *tmpItinerary)
	}

	res.Location = LocationResponse{Id: ent.Location.Id, Name: ent.Location.Name}
	if !reflect.ValueOf(ent.Airline).IsZero() {
		res.Airline = &AirlineResponse{Id: ent.Airline.Id, Name: ent.Airline.Name}
	}

	for _, rev := range ent.Reviews {
		var tmpReview = new(ReviewResponse)
		tmpReview.FromEntity(rev)

		res.Reviews = append(res.Reviews, *tmpReview)
	}
}

type ItineraryResponse struct {
	Location    string `json:"location"`
	Description string `json:"description"`
}

func (res *ItineraryResponse) FromEntity(ent tours.Itinerary) {
	res.Location = ent.Location
	res.Description = ent.Description
}

type LocationResponse struct {
	Id   uint   `json:"location_id"`
	Name string `json:"name"`
}

type AirlineResponse struct {
	Id   uint   `json:"airline_id"`
	Name string `json:"name"`
}

type ReviewResponse struct {
	User      UserResponse `json:"user"`
	Text      string       `json:"text,omitempty"`
	CreatedAt time.Time    `json:"created_at"`
}

func (res *ReviewResponse) FromEntity(ent tours.Review) {
	if !reflect.ValueOf(ent.User).IsZero() {
		var tmpUser = new(UserResponse)
		tmpUser.FromEntity(ent.User)
		res.User = *tmpUser
	}

	if ent.Text != "" {
		res.Text = ent.Text
	}

	res.CreatedAt = ent.CreatedAt
}

type UserResponse struct {
	Id    uint   `json:"user_id"`
	Name  string `json:"fullname"`
	Image string `json:"image"`
}

func (res *UserResponse) FromEntity(ent tours.User) {
	if ent.Id != 0 {
		res.Id = ent.Id
	}

	if ent.Name != "" {
		res.Name = ent.Name
	}

	if ent.Image != "" {
		res.Image = ent.Image
	} else {
		res.Image = "default"
	}
}
