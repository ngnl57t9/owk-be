package controller

import (
	"errors"
	"fmt"
	"net/http"
	"owknight-be/ent"
	"owknight-be/services/nut"
)

type AdminUserResource struct {
	*ent.Client
}

func NewAdminUserResource(client *ent.Client) *AdminUserResource {
	return &AdminUserResource{client}
}

func (res *AdminUserResource) Routes() []nut.Record {
	return []nut.Record{
		nut.Get("/", res.GetAllUsers).
			Guard(res.isAdmin).
			BindReq(&GetAllUsersReq{Pagination{0, 0}}),
	}
}

func (res *AdminUserResource) isAdmin(r *http.Request) error {
	return IsAdmin(res.Client, r)
}

type GetAllUsersReq struct {
	Pagination
}

func (res *AdminUserResource) GetAllUsers(c *nut.Context) error {
	req, ok := c.BindReq.(*GetAllUsersReq)
	if !ok {
		return errors.New("bind 失败")
	}

	u, err := res.Client.User.
		Query().
		Offset(req.Page).
		Limit(req.PageSize).
		All(c.Request.Context())
	if err != nil {
		return err
	}

	fmt.Println(u)
	return nil
}
