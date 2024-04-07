package configs

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cast"
	"reflect"
)

func unmarshal(i interface{}, o interface{}) error {
	d, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata:         nil,
		Result:           o,
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToSliceHookFunc(","),
		),
	})
	if err != nil {
		return err
	}
	return d.Decode(i)
}

func toInt64SliceE(i interface{}) ([]int64, error) {
	if i == nil {
		return nil, fmt.Errorf("unable to cast %#v of type %T to []int64", i, i)
	}

	switch v := i.(type) {
	case []int64:
		return v, nil
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]int64, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := cast.ToInt64E(s.Index(j).Interface())
			if err != nil {
				return nil, fmt.Errorf("unable to cast %#v of type %T to []int64", i, i)
			}
			a[j] = val
		}
		return a, nil
	default:
		return nil, fmt.Errorf("unable to cast %#v of type %T to []int64", i, i)
	}
}

func toFloat64SliceE(i interface{}) ([]float64, error) {
	if i == nil {
		return nil, fmt.Errorf("unable to cast %#v of type %T to []float64", i, i)
	}

	switch v := i.(type) {
	case []float64:
		return v, nil
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]float64, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := cast.ToFloat64E(s.Index(j).Interface())
			if err != nil {
				return nil, fmt.Errorf("unable to cast %#v of type %T to []float64", i, i)
			}
			a[j] = val
		}
		return a, nil
	default:
		return nil, fmt.Errorf("unable to cast %#v of type %T to []float64", i, i)
	}
}

func toStringMapFloat64E(i interface{}) (map[string]float64, error) {
	var m = map[string]float64{}
	if i == nil {
		return m, fmt.Errorf("unable to cast %#v of type %T to map[string]float64", i, i)
	}

	switch v := i.(type) {
	case map[interface{}]interface{}:
		for k, val := range v {
			m[cast.ToString(k)] = cast.ToFloat64(val)
		}
		return m, nil
	case map[string]interface{}:
		for k, val := range v {
			m[k] = cast.ToFloat64(val)
		}
		return m, nil
	case map[string]float64:
		return v, nil
	case string:
		data := []byte(v)
		return m, jsoniter.Unmarshal(data, &m)
	}

	if reflect.TypeOf(i).Kind() != reflect.Map {
		return m, fmt.Errorf("unable to cast %#v of type %T to map[string]float64", i, i)
	}
	mVal := reflect.ValueOf(m)
	v := reflect.ValueOf(i)
	for _, keyVal := range v.MapKeys() {
		val, err := cast.ToFloat64E(v.MapIndex(keyVal).Interface())
		if err != nil {
			return m, fmt.Errorf("unable to cast %#v of type %T to map[string]float64", i, i)
		}
		mVal.SetMapIndex(keyVal, reflect.ValueOf(val))
	}
	return m, nil
}
