package handler

import (
	"io"
	"wanderer/features/locations"
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
