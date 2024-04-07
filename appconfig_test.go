package configs

import (
	"context"
	"io/ioutil"
	"testing"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

func TestNewAppConfigClientErrorID(t *testing.T) {
	var err error
	_, err = newAppConfigClient(map[string]interface{}{
		"region":      "region",
		"accessKeyId": "access-key",
		"secretKey":   "secret-key",
		"app":         "sample",
		"env":         "sample",
		"configType":  "json",
		"configNames": []string{"sample"},
	})
	assert.Error(t, err)
}

func TestNewAppConfigClientErrorRegion(t *testing.T) {
	var err error
	_, err = newAppConfigClient(map[string]interface{}{
		"id":          "example-go-config-client",
		"region":      1234,
		"accessKeyId": "access-key",
		"secretKey":   "secret-key",
		"app":         "sample",
		"env":         "sample",
		"configType":  "json",
		"configNames": []string{"sample"},
	})
	assert.Error(t, err)
}

func TestNewAppConfigClientErrorAccessKeyID(t *testing.T) {
	var err error
	_, err = newAppConfigClient(map[string]interface{}{
		"id":          "example-go-config-client",
		"region":      "region",
		"secretKey":   "secret-key",
		"app":         "sample",
		"env":         "sample",
		"configType":  "json",
		"configNames": []string{"sample"},
	})
	assert.Error(t, err)
}

func TestNewAppConfigClientErrorSecretKey(t *testing.T) {
	var err error
	_, err = newAppConfigClient(map[string]interface{}{
		"id":          "example-go-config-client",
		"region":      "region",
		"accessKeyId": "access-key",
		"app":         "sample",
		"env":         "sample",
		"configType":  "json",
		"configNames": []string{"sample"},
	})
	assert.Error(t, err)
}

func TestNewAppConfigClientErrorApp(t *testing.T) {
	var err error
	_, err = newAppConfigClient(map[string]interface{}{
		"id":          "example-go-config-client",
		"region":      "region",
		"accessKeyId": "access-key",
		"secretKey":   "secret-key",
		"env":         "sample",
		"configType":  "json",
		"configNames": []string{"sample"},
	})
	assert.Error(t, err)
}

func TestNewAppConfigClientErrorEnv(t *testing.T) {
	var err error
	_, err = newAppConfigClient(map[string]interface{}{
		"id":          "example-go-config-client",
		"region":      "region",
		"accessKeyId": "access-key",
		"secretKey":   "secret-key",
		"app":         "sample",
		"configType":  "json",
		"configNames": []string{"sample"},
	})
	assert.Error(t, err)
}

func TestNewAppConfigClientErrorConfigType(t *testing.T) {
	var err error
	_, err = newAppConfigClient(map[string]interface{}{
		"id":          "example-go-config-client",
		"region":      "region",
		"accessKeyId": "access-key",
		"secretKey":   "secret-key",
		"app":         "sample",
		"env":         "sample",
		"configNames": []string{"sample"},
	})
	assert.Error(t, err)
}

func TestNewAppConfigClientErrorConfigNames(t *testing.T) {
	var err error
	_, err = newAppConfigClient(map[string]interface{}{
		"id":          "example-go-config-client",
		"region":      "region",
		"accessKeyId": "access-key",
		"secretKey":   "secret-key",
		"app":         "sample",
		"env":         "sample",
		"configType":  "json",
	})
	assert.Error(t, err)
	_, err = newAppConfigClient(map[string]interface{}{
		"id":          "example-go-config-client",
		"region":      "region",
		"accessKeyId": "access-key",
		"secretKey":   "secret-key",
		"app":         "sample",
		"env":         "sample",
		"configType":  "json",
		"configNames": 1234,
	})
	assert.Error(t, err)
}

func TestNewAppConfigClientErrorCheckInterval(t *testing.T) {
	var err error
	_, err = newAppConfigClient(map[string]interface{}{
		"id":            "example-go-config-client",
		"region":        "region",
		"accessKeyId":   "access-key",
		"secretKey":     "secret-key",
		"app":           "sample",
		"env":           "sample",
		"configType":    "json",
		"configNames":   []string{"sample"},
		"checkInterval": 1234,
	})
	assert.Error(t, err)
}

func TestNewAppConfigClientSuccess(t *testing.T) {
	var err error
	_, err = newAppConfigClient(map[string]interface{}{
		"id":                    "example-go-config-client",
		"region":                "region",
		"accessKeyId":           "access-key",
		"secretKey":             "secret-key",
		"app":                   "sample",
		"env":                   "sample",
		"configType":            "json",
		"configNames":           []string{},
		"connectTimeout":        time.Second * 10,
		"keepAliveDuration":     time.Second * 30,
		"maxIdleConnections":    100,
		"idleConnectionTimeout": time.Second * 90,
		"tlsHandshakeTimeout":   time.Second * 10,
		"expectContinueTimeout": time.Second,
		"timeout":               time.Second * 15,
	})
	assert.NoError(t, err)
}

func TestNewAppConfigClientFromAppIDSuccess(t *testing.T) {
	appID := "" // <- your app-id
	var err error
	_, err = newAppConfigClientFromAppID(context.Background(), appID, map[string]interface{}{
		"id":                    appID,
		"region":                "region",
		"accessKeyId":           "access-key",
		"secretKey":             "secret-key",
		"app":                   "app-name",
		"env":                   "app-env",
		"configType":            "yaml",
		"connectTimeout":        time.Second * 10,
		"keepAliveDuration":     time.Second * 30,
		"maxIdleConnections":    100,
		"idleConnectionTimeout": time.Second * 90,
		"tlsHandshakeTimeout":   time.Second * 10,
		"expectContinueTimeout": time.Second,
		"timeout":               time.Second * 15,
	})

	// this is to handle github test case action
	if appID == "" {
		assert.Error(t, err, "invalid appID")
		return
	}
	assert.NoError(t, err)
}

func TestAppConfigAddRemoveListenerError(t *testing.T) {
	var err error
	c, err := newAppConfigClient(map[string]interface{}{
		"id":          "example-go-config-client",
		"region":      "region",
		"accessKeyId": "access-key",
		"secretKey":   "secret-key",
		"app":         "sample",
		"env":         "sample",
		"configType":  "json",
		"configNames": []string{},
	})
	assert.NoError(t, err)
	err = c.AddChangeListener("sample", func(params ...interface{}) {})
	assert.Error(t, err)
	err = c.RemoveChangeListener("sample")
	assert.Error(t, err)
}

func TestAppConfig(t *testing.T) {
	pb, err := ioutil.ReadFile("./testresources/first.json")
	assert.NoError(t, err)
	var data map[string]interface{}
	err = jsoniter.Unmarshal(pb, &data)
	assert.NoError(t, err)
	c := &appConfigClient{
		configs: map[string]*appConfig{
			"sample": {
				data: data,
			},
		},
		listeners: make(map[string]ChangeListener),
	}
	err = c.AddChangeListener("sample", func(params ...interface{}) {})
	assert.NoError(t, err)
	err = c.RemoveChangeListener("sample")
	assert.NoError(t, err)

	g, err := c.Get("sample", "a.b.c")
	assert.NoError(t, err)
	assert.Equal(t, "d", g)
	assert.Equal(t, "naruto", c.GetD("sample", "a.b.c.s", "naruto"))

	i, err := c.GetInt("sample", "a.b.f")
	assert.NoError(t, err)
	assert.Equal(t, int64(1234), i)
	assert.Equal(t, int64(5678), c.GetIntD("sample", "a.b.c.s", 5678))

	f, err := c.GetFloat("sample", "a.b.g")
	assert.NoError(t, err)
	assert.Equal(t, 1234.5, f)
	assert.Equal(t, 5678.9, c.GetFloatD("sample", "a.b.c.s", 5678.9))

	s, err := c.GetString("sample", "a.b.c")
	assert.NoError(t, err)
	assert.Equal(t, "d", s)
	assert.Equal(t, "naruto", c.GetStringD("sample", "a.b.c.s", "naruto"))

	b, err := c.GetBool("sample", "a.b.e")
	assert.NoError(t, err)
	assert.Equal(t, true, b)
	assert.Equal(t, false, c.GetBoolD("sample", "a.b.c.s", false))

	gs, err := c.GetSlice("sample", "a.b.h")
	assert.NoError(t, err)
	assert.Equal(t, 3, len(gs))
	assert.Equal(t, 5, len(c.GetSliceD("sample", "a.b.c.s", []interface{}{1, 2, 3, 4, 5})))

	is, err := c.GetIntSlice("sample", "a.b.l")
	assert.NoError(t, err)
	assert.Equal(t, 3, len(is))
	assert.Equal(t, 5, len(c.GetIntSliceD("sample", "a.b.c.s", []int64{1, 2, 3, 4, 5})))

	fs, err := c.GetIntSlice("sample", "a.b.m")
	assert.NoError(t, err)
	assert.Equal(t, 2, len(fs))
	assert.Equal(t, 5, len(c.GetFloatSliceD("sample", "a.b.c.s", []float64{1, 2, 3, 4, 5})))

	ss, err := c.GetSlice("sample", "a.b.h")
	assert.NoError(t, err)
	assert.Equal(t, 3, len(ss))
	assert.Equal(t, 2, len(c.GetStringSliceD("sample", "a.b.c.s", []string{"naruto", "rocks"})))

	bs, err := c.GetBoolSlice("sample", "a.b.n")
	assert.NoError(t, err)
	assert.Equal(t, 3, len(bs))
	assert.Equal(t, 2, len(c.GetBoolSliceD("sample", "a.b.c.s", []bool{true, false})))

	gm, err := c.GetMap("sample", "a.b.o")
	assert.NoError(t, err)
	assert.Equal(t, 2, len(gm))
	assert.Equal(t, 1, len(c.GetMapD("sample", "a.b.c.s", map[string]interface{}{"a": "b"})))

	im, err := c.GetIntMap("sample", "a.b.v")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(im))
	assert.Equal(t, 1, len(c.GetIntMapD("sample", "a.b.c.s", map[string]int64{"a": 1})))

	fm, err := c.GetFloatMap("sample", "a.b.x")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(fm))
	assert.Equal(t, 1, len(c.GetFloatMapD("sample", "a.b.c.s", map[string]float64{"a": 1.2})))

	sm, err := c.GetStringMap("sample", "a.b.b")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(sm))
	assert.Equal(t, 1, len(c.GetStringMapD("sample", "a.b.c.s", map[string]string{"a": "b"})))

	bm, err := c.GetBoolMap("sample", "a.b.z")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(bm))
	assert.Equal(t, 1, len(c.GetBoolMapD("sample", "a.b.c.s", map[string]bool{"a": true})))
}
