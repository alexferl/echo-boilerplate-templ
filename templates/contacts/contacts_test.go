package contacts

import (
	"context"
	"io"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"

	"github.com/alexferl/echo-boilerplate-templ/models"
)

func TestTable(t *testing.T) {
	cs := models.NewContacts(1, 10)
	r, w := io.Pipe()
	go func() {
		Table(cs).Render(context.Background(), w)
		w.Close()
	}()
	doc, err := goquery.NewDocumentFromReader(r)
	assert.NoError(t, err)

	s := doc.Find(`[data-testid="contacts"]`)
	assert.True(t, s.Length() > 0)
}
