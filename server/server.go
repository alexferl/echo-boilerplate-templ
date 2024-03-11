package server

import (
	"net/http"
	"time"

	"github.com/alexferl/golib/http/api/handler"
	"github.com/alexferl/golib/http/api/middleware"
	"github.com/alexferl/golib/http/api/server"
	"github.com/labstack/echo/v4"
	_ "go.uber.org/automaxprocs"

	app "github.com/alexferl/echo-boilerplate-templ"
	"github.com/alexferl/echo-boilerplate-templ/config"
	"github.com/alexferl/echo-boilerplate-templ/handlers"
)

func New() *server.Server {
	return newServer(handlers.NewHandler())
}

func NewTestServer(handlers handlers.Handler) *server.Server {
	c := config.New()
	c.BindFlags()
	return newServer(handlers)
}

func newServer(handlers handlers.Handler) *server.Server {
	s := server.New()
	s.HTTPErrorHandler = handlers.HTTPError
	s.GET(
		"/static/*",
		handler.Static("/static/", app.StaticFS, "static"),
		middleware.Cache("/static/dist/assets/", time.Hour*24*7),
	)
	s.GET(
		"/static/images/*",
		handler.Static("/static/images/", app.StaticFS, "static/images"),
		middleware.Cache("/static/images/", time.Hour*1),
	)
	if handlers.Settings.IsProduction {
		s.GET("/static/src/*", func(c echo.Context) error {
			return c.NoContent(http.StatusForbidden)
		})
	}

	handlers.Register(s)

	return s
}
