package templates

import (
	"context"
	"io"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"

	"github.com/alexferl/echo-boilerplate-templ/models"
)

func TestError(t *testing.T) {
	r, w := io.Pipe()
	go func() {
		Error(models.HTTPError{
			Code:        "",
			Title:       "",
			Header:      "",
			Description: "",
		}).Render(context.Background(), w)
		w.Close()
	}()
	doc, err := goquery.NewDocumentFromReader(r)
	assert.NoError(t, err)

	s := doc.Find(`[data-testid="error"]`)
	assert.True(t, s.Length() > 0)
}
