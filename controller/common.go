package controller

import (
	"errors"
	"net/http"
	"owknight-be/ent"
	"owknight-be/ent/session"
)

type Pagination struct {
	Page     int `query:"page"`
	PageSize int `query:"page_size"`
}

func IsValidToken(r *http.Request) (string, error) {
	token, ok := r.Context().Value("token").(string)
	if !ok {
		return "", errors.New("token failed")
	}
	if token == "" {
		return "", errors.New("not found")
	}

	return "", nil
}

func IsAdmin(client *ent.Client, r *http.Request) error {
	token, err := IsValidToken(r)
	if err != nil {
		return err
	}

	exist, err := client.User.
		Query().
		QuerySession().
		Where(session.TokenEQ(token)).
		Exist(r.Context())
	if err != nil {
		return err
	}
	if !exist {
		return errors.New("not admin")
	}

	return nil
}
