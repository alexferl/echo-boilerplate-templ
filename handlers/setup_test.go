package handlers_test

import (
	api "github.com/alexferl/golib/http/api/server"

	"github.com/alexferl/echo-boilerplate-templ/handlers"
	"github.com/alexferl/echo-boilerplate-templ/server"
)

func getServer() *api.Server {
	h := handlers.NewHandler()
	s := server.NewTestServer(h)
	return s
}
