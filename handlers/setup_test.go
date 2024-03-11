package handlers_test

import (
	"github.com/alexferl/golib/http/api/server"

	app "github.com/alexferl/echo-boilerplate-templ"
	"github.com/alexferl/echo-boilerplate-templ/handlers"
)

func getServer() *server.Server {
	h := handlers.NewHandler()
	s := app.NewTestServer(h)
	return s
}
