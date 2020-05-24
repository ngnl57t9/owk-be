package router

import (
	"context"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"owknight-be/controller"
	"owknight-be/ent"
	"owknight-be/services/nut"
)

func V1() http.Handler {
	client := initDB()

	router := nut.NewRouter()

	router.Resource("/api/user", controller.NewUserResource(client))
	router.Resource("/api/profile", controller.NewUserResource(client))
	router.Resource("/api/session", controller.NewUserResource(client))

	router.Resource("/api/admin/user", controller.NewAdminUserResource(client))

	return router
}

func initDB() *ent.Client {
	client, err := ent.Open("postgres", "postgres://postgres:123456@localhost:5432/test?sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}
