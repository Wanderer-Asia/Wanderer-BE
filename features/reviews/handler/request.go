package handler

import "wanderer/features/reviews"

type CreateRequest struct {
	TourId uint    `json:"tour_id,omitempty"`
	Text   string  `json:"text,omitempty"`
	Rating float32 `json:"rating,omitempty"`
}

func (req *CreateRequest) ToEntity() *reviews.Review {
	var ent = new(reviews.Review)

	if req.TourId != 0 {
		ent.TourId = req.TourId
	}

	if req.Text != "" {
		ent.Text = req.Text
	}

	if req.Rating != 0 {
		ent.Rating = req.Rating
	}

	return ent
}
