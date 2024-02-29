package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/angelofallars/htmx-go"
	"github.com/labstack/echo/v4"

	"github.com/alexferl/echo-boilerplate-templ/models"
	"github.com/alexferl/echo-boilerplate-templ/templates"
	"github.com/alexferl/echo-boilerplate-templ/templates/contacts"
)

func (h *Handler) Contacts(c echo.Context) error {
	h.Settings.Title = "Contacts"

	if htmx.IsHTMX(c.Request()) {
		time.Sleep(time.Millisecond * 500) // to see the loading icon
		page := c.QueryParam("page")
		num, _ := strconv.Atoi(page)
		cs := models.NewContacts(num*10+1, num*10+10)
		return h.HTMX(h.NewHTMXResponse()).Render(c, contacts.Rows(num+1, cs))
	}

	cs := models.NewContacts(1, 10)
	return h.Render(c, http.StatusOK, templates.Base(h.Settings, contacts.Table(cs)))
}
