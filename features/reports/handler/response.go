package handler

import (
	"reflect"
	"time"
	"wanderer/features/reports"
)

type ReportResponse struct {
	TotalUser     int `json:"total_user"`
	TotalBooking  int `json:"total_booking"`
	TotalLocation int `json:"total_location"`
	TotalTour     int `json:"total_tour"`

	GraphBooking  []GraphBookingResponse `json:"graph_booking"`
	RecentBooking []BookingResponse      `json:"recent_booking"`

	TopTours []TourResponse `json:"top_tours"`
}

func (res *ReportResponse) FromEntity(ent reports.Report) {
	if ent.TotalUser != 0 {
		res.TotalUser = ent.TotalUser
	}

	if ent.TotalBooking != 0 {
		res.TotalBooking = ent.TotalBooking
	}

	if ent.TotalLocation != 0 {
		res.TotalLocation = ent.TotalLocation
	}

	if ent.TotalTour != 0 {
		res.TotalTour = ent.TotalTour
	}

	for _, graphBooking := range ent.GraphBooking {
		var tmpGraph = new(GraphBookingResponse)
		tmpGraph.FromEntity(graphBooking)
		res.GraphBooking = append(res.GraphBooking, *tmpGraph)
	}

	for _, recentBooking := range ent.RecentBooking {
		var tmpBooking = new(BookingResponse)
		tmpBooking.FromEntity(recentBooking)
		res.RecentBooking = append(res.RecentBooking, *tmpBooking)
	}

	for _, topTour := range ent.TopTours {
		var tmpTour = new(TourResponse)
		tmpTour.FromEntity(topTour)
		res.TopTours = append(res.TopTours, *tmpTour)
	}
}

type GraphBookingResponse struct {
	Label string `json:"label"`
	Total int    `json:"total"`
}

func (res *GraphBookingResponse) FromEntity(ent reports.GraphBooking) {
	if ent.Label != "" {
		res.Label = ent.Label
	}

	if ent.Total != 0 {
		res.Total = ent.Total
	}
}

type BookingResponse struct {
	Code     int     `json:"booking_code"`
	Location string  `json:"location"`
	Price    float64 `json:"price"`
}

func (res *BookingResponse) FromEntity(ent reports.Booking) {
	if ent.Code != 0 {
		res.Code = ent.Code
	}

	if ent.Location != "" {
		res.Location = ent.Location
	}

	if ent.Price != 0 {
		res.Price = ent.Price
	}
}

type TourResponse struct {
	Id       uint      `json:"tour_id"`
	Title    string    `json:"title"`
	Price    float64   `json:"price"`
	Discount int       `json:"discount"`
	Start    time.Time `json:"start"`
	Quota    int       `json:"quota"`
	Rating   float32   `json:"rating"`

	Thumbnail string `json:"thumbnail"`

	Location LocationResponse `json:"location"`
}

func (res *TourResponse) FromEntity(ent reports.Tour) {
	if ent.Id != 0 {
		res.Id = ent.Id
	}

	if ent.Title != "" {
		res.Title = ent.Title
	}

	if ent.Price != 0 {
		res.Price = ent.Price
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

	if ent.Thumbnail != "" {
		res.Thumbnail = ent.Thumbnail
	}

	if !reflect.ValueOf(ent.Location).IsZero() {
		var tmpLoc = new(LocationResponse)
		tmpLoc.FromEntity(ent.Location)

		res.Location = *tmpLoc
	}
}

type LocationResponse struct {
	Name string `json:"name"`
}

func (res *LocationResponse) FromEntity(ent reports.Location) {
	if ent.Name != "" {
		res.Name = ent.Name
	}
}
