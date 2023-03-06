package sqlboiler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitDB(t *testing.T) {
	result := SqlBiolderInitDB{}.Init(nil)
	assert.Nil(t, result)
}
