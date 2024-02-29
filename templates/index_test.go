package templates

import (
	"context"
	"io"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	r, w := io.Pipe()
	go func() {
		Index("").Render(context.Background(), w)
		w.Close()
	}()
	doc, err := goquery.NewDocumentFromReader(r)
	assert.NoError(t, err)

	s := doc.Find(`[data-testid="index"]`)
	assert.True(t, s.Length() > 0)
}
