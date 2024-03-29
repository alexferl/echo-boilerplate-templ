package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"

	_ "github.com/alexferl/echo-boilerplate-templ/testing"
)

func TestHandler_Contacts(t *testing.T) {
	s := getServer()
	req := httptest.NewRequest(http.MethodGet, "/contacts", nil)
	resp := httptest.NewRecorder()

	s.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	assert.NoError(t, err)

	sel := doc.Find(`[data-testid="contacts"]`)
	assert.True(t, sel.Length() > 0)
}

func TestHandler_Contacts_HTMX(t *testing.T) {
	s := getServer()
	req := httptest.NewRequest(http.MethodGet, "/contacts", nil)
	req.Header.Set("HX-Request", "true")
	resp := httptest.NewRecorder()

	s.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	assert.NoError(t, err)

	sel := doc.Find(`[data-testid="contacts-row"]`)
	assert.True(t, sel.Length() > 0)
}
