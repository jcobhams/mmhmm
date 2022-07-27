package config

import (
	"context"
	"fmt"
	"os"

	"github.com/jcobhams/mmhmm/logging"
)

type (
	Provider interface {
		Provide() map[string]interface{}
	}

	StaticProvider struct {
		values map[string]interface{}
	}

	StaticValue struct {
		Key   string
		Value string
	}
)

func NewStaticProvider(values ...StaticValue) *StaticProvider {
	valuesMap := make(map[string]interface{})
	for _, val := range values {
		valuesMap[val.Key] = val.Value
	}
	return &StaticProvider{values: valuesMap}
}

func (sp *StaticProvider) Provide() map[string]interface{} {
	return sp.values
}

func MustGetEnv(ctx context.Context, key string) string {
	ctx, logger := logging.Logger(ctx, "Config_MustGetEnv")
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	logger.Fatal(fmt.Sprintf("Could Not Find A Required Environment Variable: %v", key))
	return ""
}

func GetEnv(key, defaultValue string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return defaultValue
}
