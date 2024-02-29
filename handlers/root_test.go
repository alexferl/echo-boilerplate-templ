package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"

	_ "github.com/alexferl/echo-boilerplate-templ/testing"
)

func TestHandler_Root(t *testing.T) {
	s := getServer(t)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()

	s.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	assert.NoError(t, err)

	sel := doc.Find(`[data-testid="index"]`)
	assert.True(t, sel.Length() > 0)
}
