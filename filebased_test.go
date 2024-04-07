package configs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileBasedClient(t *testing.T) {
	c, err := newFileBasedClient(map[string]interface{}{
		"configsDirectory": "./testresources",
		"configNames": []string{
			"first",
			"second",
		},
		"configType": jsonType,
	})
	assert.NoError(t, err)
	assert.NotNil(t, c)
	defer func() {
		err = c.Close()
		assert.NoError(t, err)
	}()
	validateConfigData(t, c)
}

func validateConfigData(t *testing.T, c *fileBasedClient) {
	g, err := c.Get("first", "a.b.c")
	assert.NoError(t, err)
	assert.Equal(t, "d", g)
	assert.Equal(t, "naruto", c.GetD("first", "a.b.c.s", "naruto"))

	i, err := c.GetInt("first", "a.b.f")
	assert.NoError(t, err)
	assert.Equal(t, int64(1234), i)
	assert.Equal(t, int64(5678), c.GetIntD("first", "a.b.c.s", 5678))

	f, err := c.GetFloat("first", "a.b.g")
	assert.NoError(t, err)
	assert.Equal(t, 1234.5, f)
	assert.Equal(t, 5678.9, c.GetFloatD("first", "a.b.c.s", 5678.9))

	s, err := c.GetString("first", "a.b.c")
	assert.NoError(t, err)
	assert.Equal(t, "d", s)
	assert.Equal(t, "naruto", c.GetStringD("first", "a.b.c.s", "naruto"))

	b, err := c.GetBool("first", "a.b.e")
	assert.NoError(t, err)
	assert.Equal(t, true, b)
	assert.Equal(t, false, c.GetBoolD("first", "a.b.c.s", false))

	gs, err := c.GetSlice("first", "a.b.h")
	assert.NoError(t, err)
	assert.Equal(t, 3, len(gs))
	assert.Equal(t, 5, len(c.GetSliceD("first", "a.b.c.s", []interface{}{1, 2, 3, 4, 5})))

	is, err := c.GetIntSlice("first", "a.b.l")
	assert.NoError(t, err)
	assert.Equal(t, 3, len(is))
	assert.Equal(t, 5, len(c.GetIntSliceD("first", "a.b.c.s", []int64{1, 2, 3, 4, 5})))

	fs, err := c.GetIntSlice("first", "a.b.m")
	assert.NoError(t, err)
	assert.Equal(t, 2, len(fs))
	assert.Equal(t, 5, len(c.GetFloatSliceD("first", "a.b.c.s", []float64{1, 2, 3, 4, 5})))

	ss, err := c.GetSlice("first", "a.b.h")
	assert.NoError(t, err)
	assert.Equal(t, 3, len(ss))
	assert.Equal(t, 2, len(c.GetStringSliceD("first", "a.b.c.s", []string{"naruto", "rocks"})))

	bs, err := c.GetBoolSlice("first", "a.b.n")
	assert.NoError(t, err)
	assert.Equal(t, 3, len(bs))
	assert.Equal(t, 2, len(c.GetBoolSliceD("first", "a.b.c.s", []bool{true, false})))

	gm, err := c.GetMap("first", "a.b.o")
	assert.NoError(t, err)
	assert.Equal(t, 2, len(gm))
	assert.Equal(t, 1, len(c.GetMapD("first", "a.b.c.s", map[string]interface{}{"a": "b"})))

	im, err := c.GetIntMap("first", "a.b.v")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(im))
	assert.Equal(t, 1, len(c.GetIntMapD("first", "a.b.c.s", map[string]int64{"a": 1})))

	fm, err := c.GetFloatMap("first", "a.b.x")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(fm))
	assert.Equal(t, 1, len(c.GetFloatMapD("first", "a.b.c.s", map[string]float64{"a": 1.2})))

	sm, err := c.GetStringMap("first", "a.b.b")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(sm))
	assert.Equal(t, 1, len(c.GetStringMapD("first", "a.b.c.s", map[string]string{"a": "b"})))

	bm, err := c.GetBoolMap("first", "a.b.z")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(bm))
	assert.Equal(t, 1, len(c.GetBoolMapD("first", "a.b.c.s", map[string]bool{"a": true})))
}

func TestFileBasedLoaderClient(t *testing.T) {
	c, err := newDynamicFileBasedClient(map[string]interface{}{
		"configsDirectory": "./testresources",
		"configType":       yamlType,
	})
	assert.NoError(t, err)
	assert.NotNil(t, c)
	defer func() {
		err = c.Close()
		assert.NoError(t, err)
	}()
	validateConfigData(t, c)
}

func TestFileBasedLoaderClientJSON(t *testing.T) {
	c, err := newDynamicFileBasedClient(map[string]interface{}{
		"configsDirectory": "./testresources",
		"configType":       jsonType,
	})
	assert.NoError(t, err)
	assert.NotNil(t, c)
	defer func() {
		err = c.Close()
		assert.NoError(t, err)
	}()
	validateConfigData(t, c)
}
