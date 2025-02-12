package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Env        string
	Storage    Storage
	HTTPServer HTTPServer
}

type HTTPServer struct {
	Address     string
	TimeOut     time.Duration
	IdleTimeOut time.Duration
}

type Storage struct {
	Path     string
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	AsString string
}

// MustLoad loads the configuration from environment variables
func MustLoad() *Config {
	err := godotenv.Load("../../configs/env/.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	env := getEnv("ENV")
	storagePort, err := strconv.Atoi(getEnv("DB_PORT"))
	if err != nil {
		log.Fatalf("invalid DB_PORT value: %v", err)
	}

	httpTimeout, err := time.ParseDuration(getEnv("HTTP_SERVER_TIMEOUT"))
	if err != nil {
		log.Fatalf("invalid HTTP_SERVER_TIMEOUT value: %v", err)
	}

	httpIdleTimeout, err := time.ParseDuration(getEnv("HTTP_SERVER_IDLE_TIMEOUT"))
	if err != nil {
		log.Fatalf("invalid HTTP_SERVER_IDLE_TIMEOUT value: %v", err)
	}

	cfg := &Config{
		Env: env,
		Storage: Storage{
			Host:     getEnv("DB_HOST"),
			Port:     storagePort,
			User:     getEnv("DB_USER"),
			Password: getEnv("DB_PASSWORD"),
			DBName:   getEnv("DB_NAME"),
		},
		HTTPServer: HTTPServer{
			Address:     getEnv("HTTP_SERVER_ADDRESS"),
			TimeOut:     httpTimeout,
			IdleTimeOut: httpIdleTimeout,
		},
	}

	cfg.Storage.AsString = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Storage.Host, cfg.Storage.Port, cfg.Storage.User, cfg.Storage.Password, cfg.Storage.DBName)
	return cfg
}

func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	log.Fatalf("No value in Env")
	return ""
}
