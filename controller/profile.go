package controller

import "owknight-be/ent"

type ProfileResource struct {
	*ent.Client
}

func NewProfileResource(client *ent.Client) *ProfileResource {
	return &ProfileResource{client}
}

func (res *ProfileResource) GetUserProfile() error {
	return nil
}

func (res *ProfileResource) UpdateUserProfile() error {
	return nil
}
