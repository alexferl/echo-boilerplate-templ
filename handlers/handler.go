package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strconv"
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
	var m Manifest
	if isProd {
		m = LoadManifest("./static/dist/.vite/Manifest.json", "static/src/main.js")
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

func (h *Handler) Render(ctx echo.Context, statusCode int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(statusCode)
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
		text := http.StatusText(code)
		e = models.HTTPError{
			Code:   strconv.Itoa(code),
			Title:  text,
			Header: text,
		}
	}

	h.Settings.Title = e.Title
	if err := h.Render(c, code, templates.Base(h.Settings, templates.Error(e))); err != nil {
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

type Manifest struct {
	File    string   `json:"file"`
	Src     string   `json:"src"`
	IsEntry bool     `json:"isEntry"`
	CSS     []string `json:"css"`
}

func LoadManifest(path string, mainFile string) Manifest {
	plan, err := os.ReadFile(path)
	if err != nil {
		log.Panic().Err(err).Msg("failed reading manifest")
	}

	var data map[string]any
	err = json.Unmarshal(plan, &data)
	if err != nil {
		log.Panic().Err(err).Msg("failed unmarshalling manifest")
	}

	b, err := json.Marshal(data[mainFile])
	if err != nil {
		log.Panic().Err(err).Msg("failed marshalling manifest file")
	}
	var m Manifest
	err = json.Unmarshal(b, &m)
	if err != nil {
		log.Panic().Err(err).Msg("failed unmarshalling manifest file")
	}

	return m
}
