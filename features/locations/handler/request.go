package handler

import "wanderer/features/locations"

type LocationCreateUpdateRequest struct {
	Name string `json:"name"`
}

func (req *LocationCreateUpdateRequest) ToEntity() locations.Location {
	var ent = new(locations.Location)

	if req.Name != "" {
		ent.Name = req.Name
	}

	return *ent
}
