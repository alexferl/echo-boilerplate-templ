package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/a-h/templ"
	"github.com/alexferl/golib/http/api/server"
	"github.com/angelofallars/htmx-go"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/alexferl/echo-boilerplate-templ/config"
	"github.com/alexferl/echo-boilerplate-templ/models"
	"github.com/alexferl/echo-boilerplate-templ/templates"
)

type Handler struct {
	Settings models.Settings
}

func NewHandler() Handler {
	isProd := strings.ToLower(viper.GetString(config.EnvName)) == "prod"
	var m manifest
	if isProd {
		m = loadManifest()
	}
	return Handler{
		models.Settings{
			CSSFiles:     m.CSS,
			JSFile:       m.File,
			IsProduction: isProd,
			Title:        viper.GetString(config.AppName),
			ShowNav:      true,
		},
	}
}

func (h *Handler) AddRoutes(s *server.Server) {
	s.Add(http.MethodGet, "/", h.Root)
	s.Add(http.MethodGet, "/contacts", h.Contacts)
}

func (h *Handler) Render(ctx echo.Context, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(http.StatusOK)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}

func (h *Handler) HTTPError(err error, c echo.Context) {
	code := http.StatusInternalServerError
	var he *echo.HTTPError
	if errors.As(err, &he) {
		code = he.Code
	}

	e, ok := models.HTTPErrorMessages[code]
	if !ok {
		e = models.HTTPErrorMessages[http.StatusInternalServerError]
	}

	h.Settings.Title = e.Title
	if err := h.Render(c, templates.Base(h.Settings, templates.Error(e))); err != nil {
		log.Error().Err(err).Send()
	}
}

func (h *Handler) NewHTMXResponse() htmx.Response {
	return htmx.NewResponse()
}

func (h *Handler) HTMX(resp htmx.Response) HTMX {
	return NewHTMXResponse(resp)
}

type HTMX struct {
	htmx.Response
}

func NewHTMXResponse(resp htmx.Response) HTMX {
	return HTMX{resp}
}

func (h HTMX) Render(ctx echo.Context, c templ.Component) error {
	return h.RenderTempl(ctx.Request().Context(), ctx.Response(), c)
}

type manifest struct {
	File    string   `json:"file"`
	Src     string   `json:"src"`
	IsEntry bool     `json:"isEntry"`
	CSS     []string `json:"css"`
}

func loadManifest() manifest {
	plan, err := os.ReadFile("./static/dist/.vite/manifest.json")
	if err != nil {
		log.Panic().Err(err).Msg("failed reading manifest")
	}

	var data map[string]any
	err = json.Unmarshal(plan, &data)
	if err != nil {
		log.Panic().Err(err).Msg("failed unmarshalling manifest")
	}

	b, err := json.Marshal(data["static/src/main.js"])
	if err != nil {
		log.Panic().Err(err).Msg("failed marshalling manifest file")
	}
	var f manifest
	err = json.Unmarshal(b, &f)
	if err != nil {
		log.Panic().Err(err).Msg("failed unmarshalling manifest file")
	}

	return f
}
