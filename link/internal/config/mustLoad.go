package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Webhooks struct {
	Machine     string
	Shift       string
	Users       string
	Task        string
	ShiftWorker string
	ShiftTask   string
}

type Config struct {
	Env      string
	Webhooks Webhooks
	Routing  Routing
}

func MustLoad() *Config {
	err := godotenv.Load("../../configs/env/.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	env := getEnv("ENV")

	webhooks := Webhooks{
		Machine:     getEnv("WEBHOOK_MACHINE"),
		Shift:       getEnv("WEBHOOK_SHIFT"),
		Users:       getEnv("WEBHOOK_USERS"),
		Task:        getEnv("WEBHOOK_TASK"),
		ShiftWorker: getEnv("WEBHOOK_SHIFT_WORKER"),
		ShiftTask:   getEnv("WEBHOOK_SHIFT_TASK"),
	}

	cfg := &Config{
		Env:      env,
		Webhooks: webhooks,
	}

	return cfg
}

func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	log.Fatalf("No value in Env")
	return ""
}
