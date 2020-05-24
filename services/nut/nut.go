package nut

import (
	"github.com/go-chi/chi"
	"net/http"
)

func NewRouter() Router {
	root := chi.NewRouter()
	return NewMux(root, nil)
}

type Router interface {
	http.Handler

	Use(middlewares ...func(http.Handler) http.Handler)
	With(middlewares ...func(http.Handler) http.Handler) Router
	Group(fn func(r Router)) Router
	Route(pattern string, fn func(r Router)) Router
	Mount(pattern string, h http.Handler)
	Handle(pattern string, h http.Handler)
	HandleFunc(pattern string, h http.HandlerFunc)
	Method(method, pattern string, h http.Handler)
	MethodFunc(method, pattern string, h http.HandlerFunc)
	NotFound(h http.HandlerFunc)
	MethodNotAllowed(h http.HandlerFunc)

	// 修改 Handler
	Connect(pattern string, h Handler)
	Delete(pattern string, h Handler)
	Get(pattern string, h Handler)
	Head(pattern string, h Handler)
	Options(pattern string, h Handler)
	Patch(pattern string, h Handler)
	Post(pattern string, h Handler)
	Put(pattern string, h Handler)
	Trace(pattern string, h Handler)

	// 新增接口
	Resource(pattern string, ctl ResourceController)
	CatchAllError(handler CatchErrorHandler)
}

type Middlewares []func(http.Handler) http.Handler

// 新增的 ResourceController
type ResourceController interface {
	Routes() []Record
}

// 新增的 Handler
type Handler func(c *Context) error
type CatchErrorHandler func(err error, r *http.Request) error

type Guard func(r *http.Request) error
