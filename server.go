package app

import (
	"time"

	"github.com/alexferl/golib/http/api/handler"
	"github.com/alexferl/golib/http/api/middleware"
	"github.com/alexferl/golib/http/api/server"
	_ "go.uber.org/automaxprocs"

	"github.com/alexferl/echo-boilerplate-templ/config"
	"github.com/alexferl/echo-boilerplate-templ/handlers"
)

func NewServer() *server.Server {
	hs := []handlers.IHandler{
		handlers.NewHandler(),
	}
	return newServer(hs...)
}

func NewTestServer(handler ...handlers.IHandler) *server.Server {
	c := config.New()
	c.BindFlags()
	return newServer(handler...)
}

func newServer(handlers ...handlers.IHandler) *server.Server {
	s := server.New()

	s.GET(
		"/static/*",
		handler.Static("/static/", StaticFS, "static"),
		middleware.Cache("/static/dist/assets/", time.Hour*24*7),
	)
	s.GET(
		"/static/images/*",
		handler.Static("/static/images/", StaticFS, "static/images"),
		middleware.Cache("/static/images/", time.Hour*1),
	)

	for _, h := range handlers {
		h.AddRoutes(s)
	}

	return s
}
