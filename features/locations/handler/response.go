package handler

import "wanderer/features/locations"

type LocationResponse struct {
	Id    uint   `json:"location_id,omitempty"`
	Name  string `json:"name,omitempty"`
	Image string `json:"image,omitempty"`
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
}
