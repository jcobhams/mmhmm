package config

import (
	"strings"
	"sync"

	"github.com/spf13/cast"
)

type (
	Config interface {
		Get(key string) interface{}
		GetString(key string) string
		GetBytes(key string) []byte
		GetInt(key string) int
		GetFloat64(key string) float64
		GetBool(key string) bool
		FromProvider(provider Provider)
		All() map[string]interface{}
	}

	cfg struct {
		m      sync.RWMutex
		values map[string]interface{}
	}
)

func New(providers ...Provider) *cfg {
	conf := &cfg{
		values: make(map[string]interface{}),
	}

	for _, provider := range providers {
		for k, v := range provider.Provide() {
			conf.values[strings.ToLower(k)] = v
		}
	}

	return conf
}

func (c *cfg) get(key string) interface{} {
	return c.values[strings.ToLower(key)]
}

func (c *cfg) Get(key string) interface{} {
	return c.get(key)
}

func (c *cfg) GetString(key string) string {
	return cast.ToString(c.get(key))
}

func (c *cfg) GetBytes(key string) []byte {
	return []byte(cast.ToString(c.get(key)))
}

func (c *cfg) GetInt(key string) int {
	return cast.ToInt(c.get(key))
}

func (c *cfg) GetFloat64(key string) float64 {
	return cast.ToFloat64(c.get(key))
}

func (c *cfg) GetBool(key string) bool {
	return cast.ToBool(c.get(key))
}

func (c *cfg) FromProvider(provider Provider) {
	c.m.Lock()
	defer c.m.Unlock()

	for key, val := range provider.Provide() {
		c.values[strings.ToLower(key)] = val
	}
}

func (c *cfg) All() map[string]interface{} {
	return c.values
}
