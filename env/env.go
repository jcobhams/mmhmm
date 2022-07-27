package env

import (
	"context"
	"fmt"

	"github.com/jcobhams/mmhmm/config"
)

const (
	// AppEnv is the environment variable name for the application environment.
	AppEnv = "APP_ENV"

	// ServerPort is the environment variable name for the server port.
	ServerPort = "SERVER_PORT"

	// ServerAddress is the environment variable name for the server address.
	ServerAddress = "SERVER_ADDRESS"

	// DateFormat is the environment variable name for the date format.
	DateFormat = "DATE_FORMAT"
)

func ConfigProvider(ctx context.Context) *config.StaticProvider {
	defaultServerPort := "8123"

	values := []config.StaticValue{
		{Key: AppEnv, Value: config.GetEnv(AppEnv, "development")},
		{Key: ServerPort, Value: config.GetEnv(ServerPort, defaultServerPort)},
		{Key: ServerAddress, Value: fmt.Sprintf("0.0.0.0:%v", config.GetEnv(ServerPort, defaultServerPort))},
		{Key: DateFormat, Value: "2006-01-02"},
	}
	return config.NewStaticProvider(values...)
}
