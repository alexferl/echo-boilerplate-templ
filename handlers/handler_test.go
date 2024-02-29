package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/alexferl/golib/http/api/server"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/alexferl/echo-boilerplate-templ/handlers"
)

func TestLoadManifest(t *testing.T) {
	mainFile := "static/src/main.js"
	manifest := handlers.LoadManifest("./testing/fixtures/manifest.json", mainFile)
	assert.Equal(t, "assets/main-BWjwLzqI.js", manifest.File)
	assert.Equal(t, mainFile, manifest.Src)
	assert.Equal(t, []string{"assets/main-BmLLLGkb.css"}, manifest.CSS)
}

func TestHandler_HTTPError(t *testing.T) {
	testCases := []struct {
		name       string
		statusCode int
	}{
		{"403", http.StatusForbidden},
		{"404", http.StatusNotFound},
		{"410", http.StatusGone},
		{"500", http.StatusInternalServerError},
		{"502", http.StatusBadGateway},
		{"unexpected", http.StatusServiceUnavailable},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := server.New()
			handler := handlers.NewHandler()
			s.HTTPErrorHandler = handler.HTTPError
			s.GET("/error", func(c echo.Context) error { return echo.NewHTTPError(tc.statusCode) })

			req := httptest.NewRequest(http.MethodGet, "/error", nil)
			resp := httptest.NewRecorder()

			s.ServeHTTP(resp, req)
			assert.Equal(t, tc.statusCode, resp.Code)

			doc, err := goquery.NewDocumentFromReader(resp.Body)
			assert.NoError(t, err)

			sel := doc.Find(`[data-testid="error"]`)
			assert.True(t, sel.Length() > 0)
		})
	}
}
