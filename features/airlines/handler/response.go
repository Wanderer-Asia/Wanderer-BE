package handler

import "wanderer/features/airlines"

type GetAllResponse struct {
	Id    uint   `json:"airline_id,omitempty"`
	Name  string `json:"name,omitempty"`
	Image string `json:"logo,omitempty"`
}

func (res *GetAllResponse) FromEntity(ent airlines.Airline) {
	if ent.Id != 0 {
		res.Id = ent.Id
	}

	if ent.Name != "" {
		res.Name = ent.Name
	}

	if ent.ImageUrl != "" {
		res.Image = ent.ImageUrl
	} else {
		res.Image = "https://res.cloudinary.com/dhxzinjxp/image/upload/v1703490540/asset-default/plane_mefauw.png"
	}
}
