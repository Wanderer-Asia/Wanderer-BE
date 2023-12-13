package handler

import "wanderer/features/locations"

type LocationCreateRequest struct {
	Name string `json:"name"`
}

func (req *LocationCreateRequest) ToEntity() locations.Location {
	var ent = new(locations.Location)

	if req.Name != "" {
		ent.Name = req.Name
	}

	return *ent
}
