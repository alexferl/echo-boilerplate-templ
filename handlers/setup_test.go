package handlers_test

import (
	"testing"

	"github.com/alexferl/golib/http/api/server"

	app "github.com/alexferl/echo-boilerplate-templ"
	"github.com/alexferl/echo-boilerplate-templ/handlers"
)

func getServer(t *testing.T) *server.Server {
	h := handlers.NewHandler()
	s := app.NewTestServer(h)
	return s
}
