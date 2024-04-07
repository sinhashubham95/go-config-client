package configs

import (
	"context"
	"errors"
)

// These are the providers available
const (
	FileBased = iota
	AWSAppConfig
)

// generic errors
var (
	ErrConfigNotAdded = errors.New("config not added")
	ErrSecretNotAdded = errors.New("secret not added")
	ErrKeyNotFound    = errors.New("key not found")
)

// Options is the set of configurable parameters required for initialisation of the config client
type Options struct {
	// This is the provider to be used
	Provider int
	// These are the set of parameters required for the initialisation of the chosen parameter
	Params map[string]interface{}
}

// ChangeListener is called whenever any change happens in the config
type ChangeListener func(params ...interface{})

// Client is the contract that can be used and will be followed by every implementation of the config client
type Client interface {
	// AddChangeListener is used to add a listener to the changes happening to the config
	// for which it is added.
	AddChangeListener(config string, listener ChangeListener) error

	// RemoveChangeListener is used to remove the change listener added to a particular config.
	RemoveChangeListener(config string) error

	Get(config, key string) (interface{}, error)
	GetSecret(secret, key string) (interface{}, error)
	GetD(config, key string, defaultValue interface{}) interface{}
	GetInt(config, key string) (int64, error)
	GetIntSecret(secret, key string) (int64, error)
	GetIntD(config, key string, defaultValue int64) int64
	GetFloat(config, key string) (float64, error)
	GetFloatSecret(secret, key string) (float64, error)
	GetFloatD(config, key string, defaultValue float64) float64
	GetString(config, key string) (string, error)
	GetStringD(config, key string, defaultValue string) string
	GetStringSecret(secretName, key string) (string, error)
	GetBool(config, key string) (bool, error)
	GetBoolD(config, key string, defaultValue bool) bool
	GetSlice(config, key string) ([]interface{}, error)
	GetSliceD(config, key string, defaultValue []interface{}) []interface{}
	GetIntSlice(config, key string) ([]int64, error)
	GetIntSliceD(config, key string, defaultValue []int64) []int64
	GetFloatSlice(config, key string) ([]float64, error)
	GetFloatSliceD(config, key string, defaultValue []float64) []float64
	GetStringSlice(config, key string) ([]string, error)
	GetStringSliceD(config, key string, defaultValue []string) []string
	GetBoolSlice(config, key string) ([]bool, error)
	GetBoolSliceD(config, key string, defaultValue []bool) []bool
	GetMap(config, key string) (map[string]interface{}, error)
	GetMapSecret(secret, key string) (map[string]interface{}, error)
	GetMapD(config, key string, defaultValue map[string]interface{}) map[string]interface{}
	GetIntMap(config, key string) (map[string]int64, error)
	GetIntMapD(config, key string, defaultValue map[string]int64) map[string]int64
	GetFloatMap(config, key string) (map[string]float64, error)
	GetFloatMapD(config, key string, defaultValue map[string]float64) map[string]float64
	GetStringMap(config, key string) (map[string]string, error)
	GetStringMapD(config, key string, defaultValue map[string]string) map[string]string
	GetBoolMap(config, key string) (map[string]bool, error)
	GetBoolMapD(config, key string, defaultValue map[string]bool) map[string]bool

	Unmarshal(config, key string, value interface{}) error

	// Close is used to perform any closing actions on the config client
	Close() error

	ListAppConfig() []string
}

// ErrProviderNotSupported is the error used when the provider is not supported
var ErrProviderNotSupported = errors.New("provider not supported")

// New is used to initialise and get the instance of a config client
func New(options Options) (Client, error) {
	if options.Provider == FileBased {
		return newFileBasedClient(options.Params)
	}
	if options.Provider == AWSAppConfig {
		return newAppConfigClient(options.Params)
	}
	return nil, ErrProviderNotSupported
}

func NewFromAppID(configAppID string, options Options) (Client, error) {
	ctx := context.Background()
	if options.Provider == FileBased {
		return newDynamicFileBasedClient(options.Params)
	}
	if options.Provider == AWSAppConfig {
		return newAppConfigClientFromAppID(ctx, configAppID, options.Params)
	}
	return nil, ErrProviderNotSupported
}
