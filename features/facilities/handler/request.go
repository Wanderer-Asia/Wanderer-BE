package handler

import "wanderer/features/facilities"

type CreateRequest struct {
	Name string `form:"name"`
}

func (req *CreateRequest) ToEntity() *facilities.Facility {
	var ent = new(facilities.Facility)

	if req.Name != "" {
		ent.Name = req.Name
	}

	return ent
}
