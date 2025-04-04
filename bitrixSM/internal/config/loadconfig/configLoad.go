package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env        string
	WebhookB24 WebhookB24
	Webhooks   WebhookList
}

func MustLoad() *Config {
	err := godotenv.Load("../../configs/env/.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	env := getEnv("ENV")

	webhook := WebhookB24Init(
		getEnv("WEBHOOK_ID"),
		getEnv("WEBHOOK_SECRET"),
		getEnv("WEBHOOK_DOMAIN"),
		getEnv("WEBHOOK_TOKEN"),
		getEnv("WEBHOOK_URL"))
	webhooks := WebhookList{
		baseURL: getEnv("BASE_URL"),
	}
	cfg := &Config{
		Env:        env,
		WebhookB24: webhook,
		Webhooks:   webhooks,
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
