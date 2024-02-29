package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewContacts(t *testing.T) {
	contacts := NewContacts(1, 10)
	assert.Len(t, contacts, 10)
	assert.Equal(t, contacts[0].ID, "1")
	assert.Equal(t, contacts[9].ID, "10")
}
