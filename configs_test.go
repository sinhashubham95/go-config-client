package configs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	c, err := New(Options{Provider: FileBased,
		Params: map[string]interface{}{"configsDirectory": "testresources", "configNames": []string{"first"},
			"configType": jsonType}})
	assert.NoError(t, err)
	assert.NotNil(t, c)
}


func NewAppConfig(t *testing.T) {
	c, err := New(Options{Provider: AWSAppConfig,
		Params: map[string]interface{}{
			"region":      "ap-south-1",
			"accessKeyId": "yourKey",
			"secretKey":   "yourKey",
			"app":         "test",
			"env":         "dev",
			"configType":  "json",
			"configNames": []string{"test"},
			"secretNames": []string{"NiteshTest"},
			"id":          "test",
		}})
	testValue, err := c.GetStringSecret("NiteshTest", "Test")
	assert.NotNil(t, testValue)

	testIntValue, err := c.GetIntSecret("NiteshTest", "TestInt")
	assert.NotNil(t, testIntValue)

	assert.NoError(t, err)
	assert.NotNil(t, c)
}

func TestNewProviderNotSupported(t *testing.T) {
	c, err := New(Options{Provider: 1234})
	assert.Equal(t, ErrProviderNotSupported, err)
	assert.Nil(t, c)
}
