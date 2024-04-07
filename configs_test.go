package configs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	c, err := New(Options{Provider: FileBased,
		Params: map[string]interface{}{"configsDirectory": "testresources", "configNames": []string{"first"},
			"configType": jsonType}})
	assert.NoError(t, err)
	assert.NotNil(t, c)
}

func TestNewProviderNotSupported(t *testing.T) {
	c, err := New(Options{Provider: 1234})
	assert.Equal(t, ErrProviderNotSupported, err)
	assert.Nil(t, c)
}
