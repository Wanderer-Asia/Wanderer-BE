package handler

import (
	"io"
	"wanderer/features/airlines"
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
