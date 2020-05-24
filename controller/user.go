package controller

import (
	"fmt"
	"net/http"
	"owknight-be/ent"
	"owknight-be/ent/user"
	"owknight-be/services/nut"
)

type UserResource struct {
	token string
	*ent.Client
}

func NewUserResource(client *ent.Client) *UserResource {
	return &UserResource{"", client}
}

func (res *UserResource) Routes() []nut.Record {
	return []nut.Record{
		nut.Post("/register", res.Register).
			BindBody("json", &RegisterBody{}),
		nut.Post("/login", res.Login).
			BindBody("json", &LoginBody{}),
		nut.Get("/logout", res.Logout).
			Guard(res.IsValidToken),
	}
}

func (res *UserResource) IsValidToken(r *http.Request) error {
	fmt.Println(r)
	token, err := IsValidToken(r)
	res.token = token
	return err
}

type RegisterBody struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (res *UserResource) Register(c *nut.Context) error {
	reqBody := c.BindBody.(*RegisterBody)

	return c.Ok(reqBody)
}

type LoginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (res *UserResource) Login(c *nut.Context) error {
	reqBody := c.BindBody.(*LoginBody)
	fmt.Printf("body : %+v\n", reqBody)

	user, err := res.Client.User.
		Query().
		Where(user.UsernameEQ(reqBody.Username)).
		First(c.Request.Context())
	if err != nil {
		return err
	}

	// 创建 session

	cookie := http.Cookie{
		Name:  "token",
		Value: "",
	}

	// 验证用户信息是否正确
	return c.Header("cookie", cookie.String()).Ok(user)
}

func (res *UserResource) Logout(c *nut.Context) error {
	return nil
}

func (res *UserResource) checkPassword(user *ent.User) string {
	return ""
}

func (res *UserResource) genSession(session *ent.Session) string {
	return ""
}
