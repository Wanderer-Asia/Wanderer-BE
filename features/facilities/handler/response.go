package handler

import "wanderer/features/facilities"

type GetAllResponse struct {
	Id   uint   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (res *GetAllResponse) FromEntity(ent facilities.Facility) {
	if ent.Id != 0 {
		res.Id = ent.Id
	}

	if ent.Name != "" {
		res.Name = ent.Name
	}
}
