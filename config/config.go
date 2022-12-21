package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	App         string
	AppVersion  string
	Environment string //development, staging, production

	CatalogServiceGrpcHost string
	CatalogServiceGrpcPort string

	OrderServiceGrpcHost string
	OrderServiceGrpcPort string

	PostgresHost     string
	PostgresPort     int
	PostgresDatabase string
	PostgresUser     string
	PostgresPassword string

}

// Load ...
func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.App = cast.ToString(getOrReturnDefaultValue("APP", "e-commerce"))
	config.AppVersion = cast.ToString(getOrReturnDefaultValue("APP_VERSION", "1.0.0"))
	config.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", "development"))

	config.CatalogServiceGrpcHost = cast.ToString(getOrReturnDefaultValue("CATALOG_SERVICE_GRPC_HOST", "localhost"))
	config.CatalogServiceGrpcPort = cast.ToString(getOrReturnDefaultValue("CATALOG_SERVICE_GRPC_PORT", ":9001"))

	config.OrderServiceGrpcHost = cast.ToString(getOrReturnDefaultValue("ORDER_SERVICE_GRPC_HOST", "localhost"))
	config.OrderServiceGrpcPort = cast.ToString(getOrReturnDefaultValue("ORDER_SERVICE_GRPC_PORT", ":9002"))

	config.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "127.0.0.1"))
	config.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 5432))
	config.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "order_service"))
	config.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "order_service_user"))
	config.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "order_service_password"))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)

	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
