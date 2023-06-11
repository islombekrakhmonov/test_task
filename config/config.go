package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

const (
	//DebugMode indicates service mode is debug.
	DebugMode = "debug"
	// TestMode indicates service mode is test.
	TestMode = "test"
	//RealiseMode indicates service mode is release
	RealiseMode = "realise"
)

type Config struct {
	Environment string // debug, test, release

	ServerHost string
	ServerPort string

	PostgresHost     string
	PostgresUser     string
	PostgresDatabase string
	PostgresPassword string
	PostgresPort     string

	DefaultOffset int
	DefaultLimit  int

	MaxConnections int32
}


func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	cfg := Config{}

	cfg.ServerHost = "localhost"
	cfg.ServerPort = ":8080"

	cfg.PostgresHost = "localhost"
	cfg.PostgresUser = "islombek"      // установленный вами логин на postgres
	cfg.PostgresDatabase = "task"      // установленный бд на postgres
	cfg.PostgresPassword = "postgres"   // установленный вами пароль на postgres
	cfg.PostgresPort = "5432"
	cfg.DefaultOffset = 0
	cfg.DefaultLimit = 10
	cfg.MaxConnections = 100

	return cfg
}

