package templates

import (
	"context"
	"fmt"
	"io"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"

	"github.com/alexferl/echo-boilerplate-templ/models"
)

func TestBase(t *testing.T) {
	r, w := io.Pipe()
	go func() {
		Base(models.Settings{ShowNav: true}, Index("")).Render(context.Background(), w)
		w.Close()
	}()
	doc, err := goquery.NewDocumentFromReader(r)
	assert.NoError(t, err)

	for _, selector := range []string{"head", "nav", "body", "css-nonprod", "js-nonprod"} {
		s := doc.Find(fmt.Sprintf(`[data-testid="%s"]`, selector))
		assert.True(t, s.Length() > 0)

	}
}

func TestBase_NoNav(t *testing.T) {
	r, w := io.Pipe()
	go func() {
		Base(models.Settings{ShowNav: false}, Index("")).Render(context.Background(), w)
		w.Close()
	}()
	doc, err := goquery.NewDocumentFromReader(r)
	assert.NoError(t, err)

	for _, selector := range []string{"head", "body"} {
		s := doc.Find(fmt.Sprintf(`[data-testid="%s"]`, selector))
		assert.True(t, s.Length() > 0)
	}

	s := doc.Find(`[data-testid="nav"]`)
	assert.False(t, s.Length() > 0)
}

func TestBase_IsProd(t *testing.T) {
	r, w := io.Pipe()
	go func() {
		Base(models.Settings{IsProduction: true}, Index("")).Render(context.Background(), w)
		w.Close()
	}()
	doc, err := goquery.NewDocumentFromReader(r)
	assert.NoError(t, err)

	for _, selector := range []string{"css-nonprod", "js-nonprod"} {
		s := doc.Find(fmt.Sprintf(`[data-testid="%s"]`, selector))
		assert.False(t, s.Length() > 0)
	}
}
