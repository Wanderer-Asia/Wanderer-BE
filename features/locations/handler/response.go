package handler

import "wanderer/features/locations"

type LocationResponse struct {
	Id   uint   `json:"location_id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (res *LocationResponse) FromEntity(ent locations.Location) {
	if ent.Id != 0 {
		res.Id = ent.Id
	}

	if ent.Name != "" {
		res.Name = ent.Name
	}
}
