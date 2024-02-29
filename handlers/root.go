package handlers

import (
	"fmt"
	"net/http"

	libConfig "github.com/alexferl/golib/config"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"

	"github.com/alexferl/echo-boilerplate-templ/templates"
)

// Root returns the welcome message.
func (h *Handler) Root(c echo.Context) error {
	h.Settings.Title = viper.GetString(libConfig.AppName)
	msg := fmt.Sprintf("Welcome to %s!", viper.GetString(libConfig.AppName))
	return h.Render(c, http.StatusOK, templates.Base(h.Settings, templates.Index(msg)))
}
